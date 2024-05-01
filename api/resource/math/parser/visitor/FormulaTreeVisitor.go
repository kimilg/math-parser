package visitor

import parser "math-parser/parsing"

type FormulaTreeVisitor struct {
	*parser.BaseFormulaVisitor
}

func (v *FormulaTreeVisitor) VisitVariable(ctx *parser.VariableContext) interface{} {
	
	
	return v.Visit(ctx.SubscriptTail())
}