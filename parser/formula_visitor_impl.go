package parser

import (
	"github.com/antlr4-go/antlr/v4"
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
		return ctx.Constant().Accept(v)
	}
	if ctx.Variable() != nil {
		return ctx.Variable().Accept(v)
	}
	if ctx.Fraction() != nil {
		return ctx.Fraction().Accept(v)
	}
	
	if ctx.POW() != nil {
		//left := ctx.Expr(0).Accept(v)
		//right := ctx.Expr(1).Accept(v)
		
	}
	
	return ctx.GetText()
}



func (v *FormulaVisitorImpl) VisitVariable(ctx *VariableContext) interface{} {
	//name := ctx.GeneralId().Accept(v).(string)
	//subscript := ctx.SubscriptTail().Accept(v).(string)
	
	//variable := variable.DTO{Name:name, Subscripts:}
	
	//arguments := ctx.ArgumentTail().Accept(v)
	
	return nil
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

func (v *FormulaVisitorImpl) VisitGeneralId(ctx *GeneralIdContext) interface{} {
	if ctx.SINGLEID() != nil {
		return []rune(ctx.SINGLEID().GetText())
	}
	if ctx.ID() != nil {
		return []rune(ctx.ID().GetText())
	}
	return nil
}


func (v *FormulaVisitorImpl) VisitBinaryOperator(ctx *BinaryOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *FormulaVisitorImpl) VisitFraction(ctx *FractionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *FormulaVisitorImpl) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *FormulaVisitorImpl) VisitArgumentTail(ctx *ArgumentTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *FormulaVisitorImpl) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *FormulaVisitorImpl) VisitArgumentListTail(ctx *ArgumentListTailContext) interface{} {
	return v.VisitChildren(ctx)
}


func (v *FormulaVisitorImpl) VisitGeneralIntLit(ctx *GeneralIntLitContext) interface{} {
	return v.VisitChildren(ctx)
}

