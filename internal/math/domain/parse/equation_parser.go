package parse

import (
	"fmt"
	"math-parser/internal/math/domain/formula"
	"math-parser/internal/math/domain/visitor"
	"math-parser/parser"

	"github.com/antlr4-go/antlr/v4"
)

type EquationParser struct {
}

func NewEquationParser() *EquationParser {
	return &EquationParser{}
}

func (e *EquationParser) Parse(eq *formula.Equation) interface{} {
	is := antlr.NewInputStream(eq.Value)

	lexer := parser.NewFormulaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFormulaParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	formulaVisitor := visitor.FormulaVisitorImpl{Depth: 0}
	eqn := formulaVisitor.Visit(p.Equation())
	if eqn == nil {
		_ = fmt.Errorf("nil value")
	}
	exprEquation := eqn.(*formula.Expression)
	exprEquation.EquationId = eq.Id
	exprEquation.Category = eq.Category
	return exprEquation
}
