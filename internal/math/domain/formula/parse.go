package formula

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"math-parser/parser"
)

func ParseEquation(eq *Equation) *Expression {
	is := antlr.NewInputStream(eq.Value)

	lexer := parser.NewFormulaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFormulaParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	formulaVisitor := FormulaVisitorImpl{depth: 0}
	eqn := formulaVisitor.Visit(p.Equation())
	if eqn == nil {
		fmt.Errorf("nil value")
	}
	exprEquation := eqn.(*Expression)
	exprEquation.EquationId = eq.Id
	exprEquation.Category = eq.Category
	return exprEquation
}