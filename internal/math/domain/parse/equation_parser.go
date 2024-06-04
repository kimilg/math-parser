package parse

import (
	"fmt"
	"math-parser/internal/math/domain/formula"
	"math-parser/parser"

	"github.com/antlr4-go/antlr/v4"
)

type EquationParser struct {
	visitor parser.FormulaVisitor
}

func NewEquationParser(visitor parser.FormulaVisitor) *EquationParser {
	return &EquationParser{
		visitor: visitor,
	}
}

func (e *EquationParser) Parse(eq *formula.Equation) interface{} {
	is := antlr.NewInputStream(eq.Value)

	lexer := parser.NewFormulaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFormulaParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	eqn := e.visitor.Visit(p.Equation())
	if eqn == nil {
		_ = fmt.Errorf("nil value")
	}
	exprEquation := eqn.(*formula.Expression)
	exprEquation.EquationId = eq.Id
	exprEquation.Category = eq.Category
	return exprEquation
}
