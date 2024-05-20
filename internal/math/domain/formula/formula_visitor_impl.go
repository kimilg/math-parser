package formula

import (
	"fmt"
	"math-parser/parser"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type FormulaVisitorImpl struct {
	*parser.BaseFormulaVisitor
	depth int
}

func (v *FormulaVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.EquationContext:
		return val.Accept(v)
	}
	return nil
}

func (v *FormulaVisitorImpl) VisitEquation(ctx *parser.EquationContext) interface{} {
	return &Expression{
		Elements: []Element{
			ctx.Expr(0).Accept(v),
			EQUAL,
			ctx.Expr(1).Accept(v),
		}}
}

func (v *FormulaVisitorImpl) returnExprOrArgument(expr *Expression) interface{} {
	if v.depth == 0 {
		return expr
	}
	return &Argument{Expression: expr}
}

func (v *FormulaVisitorImpl) VisitExpr(ctx *parser.ExprContext) interface{} {
	str := ctx.GetText()
	println(str)
	if ctx.OPENPAREN() != nil {
		return ctx.Expr(0).Accept(v)
	}
	if ctx.Constant() != nil {
		return &Expression{
			Elements: []Element{
				&Constant{
					Value: ctx.Constant().Accept(v).(float64),
				},
			},
			Description: "constant",
			IsArgument:  v.depth == 0,
		}
	}
	if ctx.Variable() != nil {
		return &Expression{
			Elements: []Element{
				ctx.Variable().Accept(v).(Variable),
			},
			Description: "variable",
			IsArgument:  v.depth == 0,
		}
	}
	if ctx.Fraction() != nil {
		return ctx.Fraction().Accept(v)
	}
	if ctx.POW() != nil {
		return v.exprForBinaryOp(ctx, POW)
	}
	if ctx.MULT() != nil {
		return v.exprForBinaryOp(ctx, MULT)
	}
	if ctx.DIV() != nil {
		return v.exprForBinaryOp(ctx, DIV)
	}
	if ctx.PLUS() != nil {
		return v.exprForBinaryOp(ctx, PLUS)
	}
	if ctx.MINUS() != nil {
		return v.exprForBinaryOp(ctx, MINUS)
	}

	return nil
}

func (v *FormulaVisitorImpl) VisitVariable(ctx *parser.VariableContext) interface{} {
	str := ctx.GetText()
	println(str)

	name := ctx.GeneralId().Accept(v).(string)
	var subscripts []rune
	var arguments []*Argument

	if ctx.SubscriptTail() != nil && ctx.SubscriptTail().SUBSCRIPT() != nil {
		str = ctx.SubscriptTail().GetText()
		subscripts = []rune(ctx.SubscriptTail().Accept(v).(string))
	}
	if ctx.ArgumentTail() != nil && ctx.ArgumentTail().OPENPAREN() != nil {
		str = ctx.ArgumentTail().GetText()
		v.depth += 1
		arguments = ctx.ArgumentTail().Accept(v).([]*Argument)
		v.depth -= 1
	}

	variable := Variable{
		Name:        name,
		Subscripts:  subscripts,
		Arguments:   arguments,
		Description: "description..",
	}

	return variable
}

func (v *FormulaVisitorImpl) VisitSubscriptTail(ctx *parser.SubscriptTailContext) interface{} {
	if ctx.SINGLEID() != nil {
		return ctx.SINGLEID().GetText()
	}
	if ctx.SINGLEINTLIT() != nil {
		return ctx.SINGLEINTLIT().GetText()
	}
	if ctx.GeneralId() != nil {
		return ctx.GeneralId().Accept(v).(string)
	}
	if ctx.GeneralIntLit() != nil {
		return strconv.Itoa(ctx.GeneralIntLit().Accept(v).(int))
	}

	return nil
}

func (v *FormulaVisitorImpl) VisitArgumentTail(ctx *parser.ArgumentTailContext) interface{} {
	if ctx.SEMICOLON() == nil {
		return ctx.ArgumentList(0).Accept(v).([]*Argument)
	}

	str := ctx.ArgumentList(0).GetText()
	println(str)
	left := ctx.ArgumentList(0).Accept(v).([]*Argument)
	for _, argument := range left {
		argument.IsEffect = true
	}

	str = ctx.ArgumentList(1).GetText()
	println(str)
	right := ctx.ArgumentList(1).Accept(v).([]*Argument)
	for _, argument := range right {
		argument.IsCause = true
	}

	arguments := left
	arguments = append(arguments, right...)
	return arguments
}

func (v *FormulaVisitorImpl) VisitGeneralId(ctx *parser.GeneralIdContext) interface{} {
	if ctx.SINGLEID() != nil {
		return ctx.SINGLEID().GetText()
	}
	if ctx.ID() != nil {
		return ctx.ID().GetText()
	}
	return nil
}

func (v *FormulaVisitorImpl) VisitConstant(ctx *parser.ConstantContext) interface{} {
	if ctx.GeneralIntLit() != nil {
		return float64(ctx.GeneralIntLit().Accept(v).(int))
	}
	if ctx.FLOATLIT() != nil {
		return ctx.FLOATLIT().Accept(v).(float64)
	}

	return nil
}

func (v *FormulaVisitorImpl) VisitFraction(ctx *parser.FractionContext) interface{} {
	return &Expression{
		Elements: []Element{
			ctx.Expr(0).Accept(v),
			&Operator{Value: DIV},
			ctx.Expr(1).Accept(v),
		},
		IsArgument: v.depth == 0,
	}
}

func (v *FormulaVisitorImpl) VisitBinaryOperator(ctx *parser.BinaryOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *FormulaVisitorImpl) VisitArgumentList(ctx *parser.ArgumentListContext) interface{} {
	arguments := []*Argument{{Expression: ctx.Expr().Accept(v).(*Expression)}}
	if ctx.ArgumentListTail() != nil {
		arguments = append(arguments, ctx.ArgumentListTail().Accept(v).([]*Argument)...)
	}
	return arguments
}

func (v *FormulaVisitorImpl) VisitArgumentListTail(ctx *parser.ArgumentListTailContext) interface{} {
	arguments := []*Argument{{Expression: ctx.Expr().Accept(v).(*Expression)}}
	if ctx.ArgumentListTail() != nil && ctx.ArgumentListTail().COMMA() != nil {
		arguments = append(arguments, ctx.ArgumentListTail().Accept(v).([]*Argument)...)
	}
	return arguments
}

func (v *FormulaVisitorImpl) VisitGeneralIntLit(ctx *parser.GeneralIntLitContext) interface{} {
	if ctx.SINGLEINTLIT() != nil {
		singleIntLit, err := strconv.Atoi(ctx.SINGLEINTLIT().GetText())
		if err != nil {
			fmt.Errorf("singleIntLit %s should be integer: %w", ctx.SINGLEINTLIT().GetText(), err)
		}
		return singleIntLit
	}
	if ctx.INTLIT() != nil {
		intLit, err := strconv.Atoi(ctx.INTLIT().GetText())
		if err != nil {
			fmt.Errorf("intLit %s should be integer: %w", ctx.INTLIT().GetText(), err)
		}
		return intLit
	}

	return nil
}

func (v *FormulaVisitorImpl) exprForBinaryOp(ctx *parser.ExprContext, operator byte) *Expression {
	return &Expression{
		Elements: []Element{
			ctx.Expr(0).Accept(v),
			operator,
			ctx.Expr(1).Accept(v),
		},
		IsArgument: v.depth == 0,
	}
}
