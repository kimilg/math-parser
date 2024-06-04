package visitor

import (
	"math-parser/parser"

	"github.com/antlr4-go/antlr/v4"
)

type FormulaListenerImpl struct {
	*parser.BaseFormulaListener
}

func (l *FormulaListenerImpl) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//println(ctx.GetText())
}

func (l *FormulaListenerImpl) EnterVariable(ctx *parser.VariableContext) {
	println(ctx.GetText())
}

func (l *FormulaListenerImpl) EnterConstant(ctx *parser.ConstantContext) {
	println(ctx.GetText())
}

func (l *FormulaListenerImpl) EnterGeneralId(ctx *parser.GeneralIdContext) {
	println(ctx.GetText())
}
