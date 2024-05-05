package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"math-parser/api/resource/math"
	"strconv"
)

type FormulaVisitorImpl struct {
	*BaseFormulaVisitor
}

func (v *FormulaVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *EquationContext:
		return val.Accept(v)
	}
	return nil
}

func (v *FormulaVisitorImpl) VisitEquation(ctx *EquationContext) interface{} {
	left := ctx.Expr(0).Accept(v).(string) // new expr obj ?
	right := ctx.Expr(1).Accept(v).(string)
	return left + right
}

func (v *FormulaVisitorImpl) VisitExpr(ctx *ExprContext) interface{} {
	if ctx.OPENPAREN() != nil {
		return ctx.Expr(0).Accept(v)
	}
	if ctx.Constant() != nil {
		return math.Expression{
			Elements: []math.Element{
				math.Constant{
					Value: ctx.Constant().Accept(v).(float64),
				},
			},
			Description: "constant",
		}
	}
	if ctx.Variable() != nil {
		return math.Expression{
			Elements: []math.Element{
				ctx.Variable().Accept(v).(math.Variable),
			},
			Description: "variable",
		}
	}
	if ctx.Fraction() != nil {
		return ctx.Fraction().Accept(v)
	}
	if ctx.POW() != nil {
		return v.exprForBinaryOp(ctx, math.POW)
	}
	if ctx.MULT() != nil {
		return v.exprForBinaryOp(ctx, math.MULT)
	}
	if ctx.DIV() != nil {
		return v.exprForBinaryOp(ctx, math.DIV)
	}
	if ctx.PLUS() != nil {
		return v.exprForBinaryOp(ctx, math.PLUS)
	}
	if ctx.MINUS() != nil {
		return v.exprForBinaryOp(ctx, math.MINUS)
	}
	
	return nil
}

func (v *FormulaVisitorImpl) VisitVariable(ctx *VariableContext) interface{} {
	name := ctx.GeneralId().Accept(v).(string)
	subscripts := []rune(ctx.SubscriptTail().Accept(v).(string))
	arguments := ctx.ArgumentTail().Accept(v).([]math.Expression)
	
	variable := math.Variable{
		Name:        name,
		Subscripts:  subscripts,
		Arguments:   arguments,
		Description: "description..",
	}
	
	return variable
}

func (v *FormulaVisitorImpl) VisitSubscriptTail(ctx *SubscriptTailContext) interface{} {
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
		return ctx.GeneralIntLit().Accept(v).(string)
	}
	
	return nil
}

func (v *FormulaVisitorImpl) VisitArgumentTail(ctx *ArgumentTailContext) interface{} {
	if ctx.SEMICOLON() == nil {
		return ctx.ArgumentList(0).Accept(v).([]math.Expression)
	}
	
	left := ctx.ArgumentList(0).Accept(v).([]math.Expression)
	for _, expr := range left {
		expr.IsEffect = true
	}

	right := ctx.ArgumentList(1).Accept(v).([]math.Expression)
	for _, expr := range right {
		expr.IsCause = true
	}
	
	expressions := left
	expressions = append(expressions, right...)
	return expressions
}


func (v *FormulaVisitorImpl) VisitGeneralId(ctx *GeneralIdContext) interface{} {
	if ctx.SINGLEID() != nil {
		return ctx.SINGLEID().GetText()
	}
	if ctx.ID() != nil {
		return ctx.ID().GetText()
	}
	return nil
}

func (v *FormulaVisitorImpl) VisitConstant(ctx *ConstantContext) interface{} {
	if ctx.GeneralIntLit() != nil {
		return ctx.GeneralIntLit().Accept(v).(float64)
	}
	if ctx.FLOATLIT() != nil {
		return ctx.FLOATLIT().Accept(v).(float64)
	}
	
	return nil
}

func (v *FormulaVisitorImpl) VisitFraction(ctx *FractionContext) interface{} {
	return math.Expression{Elements: []math.Element{
		ctx.Expr(0).Accept(v),
		math.Operator{Value: math.DIV},
		ctx.Expr(1).Accept(v),
	}}
}


func (v *FormulaVisitorImpl) VisitBinaryOperator(ctx *BinaryOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *FormulaVisitorImpl) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	expressions := []math.Expression{ctx.Expr().Accept(v).(math.Expression)}
	if ctx.ArgumentListTail() != nil {
		expressions = append(expressions, ctx.ArgumentListTail().Accept(v).([]math.Expression)...)
	}
	return expressions
}

func (v *FormulaVisitorImpl) VisitArgumentListTail(ctx *ArgumentListTailContext) interface{} {
	expressions := []math.Expression{ctx.Expr().Accept(v).(math.Expression)}
	if ctx.ArgumentListTail() != nil {
		expressions = append(expressions, ctx.ArgumentListTail().Accept(v).([]math.Expression)...)
	}
	return expressions
}

func (v *FormulaVisitorImpl) VisitGeneralIntLit(ctx *GeneralIntLitContext) interface{} {
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

func (v *FormulaVisitorImpl) exprForBinaryOp(ctx *ExprContext, operator byte) math.Expression {
	return math.Expression{
		Elements: []math.Element{
			ctx.Expr(0).Accept(v),
			operator,
			ctx.Expr(1).Accept(v),
		}}
}
