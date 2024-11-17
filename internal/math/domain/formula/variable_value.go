package formula

import "math-parser/internal/math/domain/field"

type VariableValue struct {
	Equation  string            `json:"equation" form:"required"`
	Variables []*field.Variable `json:"variables" form:"required"`
}
