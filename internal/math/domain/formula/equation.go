package formula

import (
	"context"
	"fmt"
)

type ID int64

type Equation struct {
	Id    ID     `json:"id" form:"required"`
	Value string `json:"equation" form:"required"`
}

type EquationMemory struct {
	equations map[ID]*Expression
}

func NewEquationMemory() EquationMemory {
	return EquationMemory{make(map[ID]*Expression)}
}

func (e EquationMemory) Get(ctx context.Context, id ID) (*Expression, error) {
	if e.equations[id] == nil {
		return nil, fmt.Errorf("no such id: %d", id)
	}
	return e.equations[id], nil
}

func (e EquationMemory) List(ctx context.Context) ([]*Expression, error) {
	var expressions []*Expression
	for _, v := range e.equations {
		expressions = append(expressions, v)
	}
	return expressions, nil
}
