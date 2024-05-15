package formula

import (
	"context"
	"fmt"
)

type ID int64

type Equation struct {
	Id    ID     `json:"id" form:"required"`
	Value string `json:"value" form:"required"`
	Category string `json:"category" form:"required"`
	Variables []Variable `json:"variables"`
	Cause string `json:"cause"`
	Effect string `json:"effect"`
}

type EquationMemory struct {
	equations map[ID]*Expression
	repo Repository
}

func NewEquationMemory(repo Repository) EquationMemory {
	return EquationMemory{
		make(map[ID]*Expression),
		repo,
	}
}

func (e *EquationMemory) Load(ctx context.Context) error {
	equations, err := e.repo.List(ctx)
	if err != nil {
		return fmt.Errorf("error during loading equation repository: %v", err)
	}
	for _, equation := range equations {
		expression := ParseEquation(equation)
		e.equations[expression.EquationId] = expression
	}
	return nil
}

func (e *EquationMemory) Get(id ID) (*Expression, error) {
	if e.equations[id] == nil {
		return nil, fmt.Errorf("no such id: %d", id)
	}
	return e.equations[id], nil
}

func (e *EquationMemory) List() ([]*Expression, error) {
	var expressions []*Expression
	for _, v := range e.equations {
		expressions = append(expressions, v)
	}
	return expressions, nil
}
