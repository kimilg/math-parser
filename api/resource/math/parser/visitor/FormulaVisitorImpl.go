package visitor

import (
	"github.com/antlr4-go/antlr/v4"
	parser "math-parser/parsing"
)

type FormulaVisitorImpl struct {
	*parser.BaseFormulaVisitor
}

func (v *FormulaVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.EquationContext:
		return val.Accept(v) //여기서 FormulaVisitor 타입이 나오지 않는다.
	}
	return nil
}

func (v *FormulaVisitorImpl) VisitEquation(ctx *parser.EquationContext) interface{} {
	left := ctx.Expr(0).Accept(v).(string) // new expr obj ?
	right := ctx.Expr(1).Accept(v).(string)
	return left + right
}

func (v *FormulaVisitorImpl) VisitVariable(ctx *parser.VariableContext) interface{} {

	ctx.SubscriptTail().Accept(v)
	
	return v.Visit(ctx.SubscriptTail())
}

func (v *FormulaVisitorImpl) VisitSubscriptTail(ctx *parser.SubscriptTailContext) interface{} {
	subscript := ctx.SUBSCRIPT()
	openCurly := ctx.OPENCURLY()
	generalId := ctx.GeneralId()
	closeCurly := ctx.CLOSECURLY()
	generalIntLit := ctx.GeneralIntLit()
	singleId := ctx.SINGLEID()
	singleIntLit := ctx.SINGLEINTLIT()
	
	if subscript == nil || openCurly == nil || generalId == nil || closeCurly == nil || generalIntLit == nil ||
		singleId == nil || singleIntLit == nil {
		println("yes.")
	}
	
	return ctx.GeneralId().Accept(v)
}

func (v *FormulaVisitorImpl) VisitExpr(ctx *parser.ExprContext) interface{} {
	if ctx.Variable() != nil {
		return ctx.Variable().Accept(v)
	}
	return ctx.GetText()
}

func (v *FormulaVisitorImpl) VisitBinaryOperator(ctx *parser.BinaryOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *FormulaVisitorImpl) VisitFraction(ctx *parser.FractionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *FormulaVisitorImpl) VisitConstant(ctx *parser.ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *FormulaVisitorImpl) VisitArgumentTail(ctx *parser.ArgumentTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *FormulaVisitorImpl) VisitArgumentList(ctx *parser.ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *FormulaVisitorImpl) VisitArgumentListTail(ctx *parser.ArgumentListTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *FormulaVisitorImpl) VisitGeneralId(ctx *parser.GeneralIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *FormulaVisitorImpl) VisitGeneralIntLit(ctx *parser.GeneralIntLitContext) interface{} {
	return v.VisitChildren(ctx)
}

