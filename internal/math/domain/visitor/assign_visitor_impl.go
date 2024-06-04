package visitor

import (
	"math-parser/parser"

	"github.com/antlr4-go/antlr/v4"
)

type AssignVisitorImpl struct {
	*parser.BaseFormulaVisitor
}

func (v *AssignVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.EquationContext:
		return val.Accept(v)
	}
	return nil
}

func (v *AssignVisitorImpl) VisitEquation(ctx *parser.EquationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AssignVisitorImpl) VisitExpr(ctx *parser.ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AssignVisitorImpl) VisitBinaryOperator(ctx *parser.BinaryOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AssignVisitorImpl) VisitFraction(ctx *parser.FractionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AssignVisitorImpl) VisitConstant(ctx *parser.ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AssignVisitorImpl) VisitVariable(ctx *parser.VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AssignVisitorImpl) VisitSubscriptTail(ctx *parser.SubscriptTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AssignVisitorImpl) VisitArgumentTail(ctx *parser.ArgumentTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AssignVisitorImpl) VisitArgumentList(ctx *parser.ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AssignVisitorImpl) VisitArgumentListTail(ctx *parser.ArgumentListTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AssignVisitorImpl) VisitGeneralId(ctx *parser.GeneralIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AssignVisitorImpl) VisitGeneralIntLit(ctx *parser.GeneralIntLitContext) interface{} {
	return v.VisitChildren(ctx)
}
