package parser

import (
	"github.com/antlr4-go/antlr/v4"
)

type FormulaListenerImpl struct {
	*BaseFormulaListener
}

func (l *FormulaListenerImpl) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//println(ctx.GetText())
}

func (l *FormulaListenerImpl) EnterVariable(ctx *VariableContext) {
	println(ctx.GetText())
}

func (l *FormulaListenerImpl) EnterConstant(ctx *ConstantContext) {
	println(ctx.GetText())
}

func (l *FormulaListenerImpl) EnterGeneralId(ctx *GeneralIdContext) {
	println(ctx.GetText())
}
