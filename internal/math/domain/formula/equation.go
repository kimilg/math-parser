package formula

import (
	"context"
	"fmt"
)

type ID int64

type Equation struct {
	Id        ID          `json:"id" form:"required"`
	Value     string      `json:"value" form:"required"`
	Category  string      `json:"category" form:"required"`
	Variables []*Variable `json:"variables"`
	Cause     string      `json:"cause"`
	Effect    string      `json:"effect"`
}

type EquationMemory struct {
	equations   map[ID]*Equation
	expressions map[ID]*Expression
	repo        Repository
	parser      Parser
}

type Pair struct {
	Equation *Equation
	Expression *Expression
}

func NewEquationMemory(repo Repository, parser Parser) *EquationMemory {
	return &EquationMemory{
		make(map[ID]*Equation),
		make(map[ID]*Expression),
		repo,
		parser,
	}
}

func (e *EquationMemory) Load(ctx context.Context) error {
	equations, err := e.repo.List(ctx)
	if err != nil {
		return fmt.Errorf("error during loading equation repository: %v", err)
	}
	for _, equation := range equations {
		e.equations[equation.Id] = equation
		expression := e.parser.Parse(equation).(*Expression)
		e.expressions[expression.EquationId] = expression
	}
	return nil
}

func (e *EquationMemory) GetEquation(id ID) (*Equation, error) {
	if e.equations[id] == nil {
		return nil, fmt.Errorf("no such id: %d", id)
	}
	return e.equations[id], nil
}

func (e *EquationMemory) GetExpression(id ID) (*Expression, error) {
	if e.expressions[id] == nil {
		return nil, fmt.Errorf("no such id: %d", id)
	}
	return e.expressions[id], nil
}

func (e *EquationMemory) List() ([]*Pair, error) {
	var pairs []*Pair
	for k, v := range e.equations {
		pairs = append(pairs, &Pair{v, e.expressions[k]})
	}
	return pairs, nil
}

func (e *EquationMemory) ListEquations() ([]*Equation, error) {
	var equations []*Equation
	for _, v := range e.equations {
		equations = append(equations, v)
	}
	return equations, nil
}

func (e *EquationMemory) ListExpressions() ([]*Expression, error) {
	var expressions []*Expression
	for _, v := range e.expressions {
		expressions = append(expressions, v)
	}
	return expressions, nil
}

func (e *EquationMemory) Insert(equation *Equation) (*Equation, error) {
	if equation == nil {
		return nil, fmt.Errorf("expression is nil")
	}
	e.equations[equation.Id] = equation
	expression := e.parser.Parse(equation).(*Expression)
	e.expressions[expression.EquationId] = expression
	return equation, nil
}

func (e *EquationMemory) InsertEquation(equation *Equation) (*Equation, error) {
	if equation == nil {
		return nil, fmt.Errorf("expression is nil")
	}
	e.equations[equation.Id] = equation
	return equation, nil
}

func (e *EquationMemory) InsertExpression(expression *Expression) (*Expression, error) {
	if expression == nil {
		return nil, fmt.Errorf("expression is nil")
	}
	e.expressions[expression.EquationId] = expression
	return expression, nil
}
