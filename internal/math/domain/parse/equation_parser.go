package parse

import (
	"fmt"
	"math-parser/internal/math/domain/formula"
	"math-parser/parser"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

type EquationParser struct {
	visitor parser.FormulaVisitor
	ParserMapper map[string]*parser.FormulaParser
	mutex *sync.Mutex
}

func NewEquationParser(visitor parser.FormulaVisitor) *EquationParser {
	return &EquationParser{
		visitor: visitor,
		mutex:   &sync.Mutex{},
	}
}

func (e *EquationParser) Parse(eq *formula.Equation) interface{} {
	is := antlr.NewInputStream(eq.Value)

	lexer := parser.NewFormulaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFormulaParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	expr := e.visitor.Visit(p.Equation())
	if expr == nil {
		_ = fmt.Errorf("nil value")
	}
	exprEquation := expr.(*formula.Expression)
	exprEquation.EquationId = eq.Id
	exprEquation.Category = eq.Category
	return exprEquation
}

func (e *EquationParser) GetParser(eqStr string) *parser.FormulaParser {

	if e.ParserMapper[eqStr] == nil {
		e.mutex.Lock()
		defer e.mutex.Unlock()
		if e.ParserMapper[eqStr] == nil {
			is := antlr.NewInputStream(eqStr)

			lexer := parser.NewFormulaLexer(is)
			stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

			p := parser.NewFormulaParser(stream)
			p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

			e.ParserMapper[eqStr] = p
		}
	}

	return e.ParserMapper[eqStr]
}
