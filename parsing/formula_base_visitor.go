// Code generated from Formula.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Formula
import "github.com/antlr4-go/antlr/v4"

type BaseFormulaVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseFormulaVisitor) VisitEquation(ctx *EquationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaVisitor) VisitBinaryOperator(ctx *BinaryOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaVisitor) VisitFraction(ctx *FractionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaVisitor) VisitSubscriptTail(ctx *SubscriptTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaVisitor) VisitArgumentTail(ctx *ArgumentTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaVisitor) VisitArgumentListTail(ctx *ArgumentListTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaVisitor) VisitGeneralId(ctx *GeneralIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaVisitor) VisitGeneralIntLit(ctx *GeneralIntLitContext) interface{} {
	return v.VisitChildren(ctx)
}
