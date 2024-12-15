package entity

import (
	"math-parser/internal/math/domain/field"
	"math-parser/internal/math/domain/formula"
)

type MathInfo struct {
	Equation       string
	Expression     *formula.Expression
	VariableMapper *field.VariableMapper
	ArgumentKinds  []*formula.ArgumentKind
	ArgumentMapper map[string]field.IVector
}
