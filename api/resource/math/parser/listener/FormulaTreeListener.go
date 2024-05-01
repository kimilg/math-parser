package listener

import (
	"github.com/antlr4-go/antlr/v4"
	parser "math-parser/parsing"
)

type FormulaTreeListener struct {
	*parser.BaseFormulaListener
}

func (l *FormulaTreeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//println(ctx.GetText())
}

func (l *FormulaTreeListener) EnterVariable(ctx *parser.VariableContext) {
	println(ctx.GetText())
}

func (l *FormulaTreeListener) EnterConstant(ctx *parser.ConstantContext) {
	println(ctx.GetText())
}
