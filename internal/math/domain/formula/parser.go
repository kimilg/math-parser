package formula

import "math-parser/parser"

type Parser interface {
	Parse(eq *Equation) interface{}
	GetParser(eqStr string) *parser.FormulaParser
}
