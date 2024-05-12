package formula

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"math-parser/parser"
)

func ParseEquation(id ID, value string) *Expression {
	is := antlr.NewInputStream(value)

	lexer := parser.NewFormulaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFormulaParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	var formulaVisitor FormulaVisitorImpl
	eqn := formulaVisitor.Visit(p.Equation())
	if eqn == nil {
		fmt.Errorf("nil value")
	}
	exprEquation := eqn.(*Expression)
	exprEquation.EquationId = id
	return exprEquation
}