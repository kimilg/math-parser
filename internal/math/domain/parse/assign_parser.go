package parse

import (
	"fmt"
	"math-parser/internal/math/domain/formula"
	"math-parser/parser"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

type AssignParser struct {
	visitor parser.FormulaVisitor
	ParserMapper map[string]*parser.FormulaParser
	mutex *sync.Mutex
}

func NewAssignParser(visitor parser.FormulaVisitor) *AssignParser {
	return &AssignParser{
		visitor: visitor,
		ParserMapper: make(map[string]*parser.FormulaParser),
		mutex: &sync.Mutex{},
	}
}

func (a *AssignParser) Parse(eq *formula.Equation) interface{} {
	is := antlr.NewInputStream(eq.Value)

	lexer := parser.NewFormulaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFormulaParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	eqn := a.visitor.Visit(p.Equation())
	if eqn == nil {
		_ = fmt.Errorf("nil value")
	}
	exprEquation := eqn.(*formula.Expression)
	exprEquation.EquationId = eq.Id
	exprEquation.Category = eq.Category
	return exprEquation
}

func (a *AssignParser) GetParser(eqStr string) *parser.FormulaParser {
	
	if a.ParserMapper[eqStr] == nil {
		a.mutex.Lock()
		defer a.mutex.Unlock()
		if a.ParserMapper[eqStr] == nil {
			is := antlr.NewInputStream(eqStr)

			lexer := parser.NewFormulaLexer(is)
			stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

			p := parser.NewFormulaParser(stream)
			p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

			a.ParserMapper[eqStr] = p
		}
	}
	
	return a.ParserMapper[eqStr]
}
