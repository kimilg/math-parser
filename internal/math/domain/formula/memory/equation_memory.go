package memory

import (
	"context"
	"fmt"
	"math-parser/internal/math/domain/formula"
)

type EquationMemory struct {
	equations         map[formula.ID]*formula.Equation
	equationFromValue map[string]*formula.Equation
	expressions       map[formula.ID]*formula.Expression
	repo              formula.Repository
	parser            formula.Parser
}

type Pair struct {
	Equation   *formula.Equation
	Expression *formula.Expression
}

func NewEquationMemory(repo formula.Repository, parser formula.Parser) *EquationMemory {
	return &EquationMemory{
		make(map[formula.ID]*formula.Equation),
		make(map[string]*formula.Equation),
		make(map[formula.ID]*formula.Expression),
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
		e.equationFromValue[equation.Value] = equation
		expression := e.parser.Parse(equation).(*formula.Expression)
		e.expressions[expression.EquationId] = expression
	}
	return nil
}

func (e *EquationMemory) GetEquation(id formula.ID) (*formula.Equation, error) {
	if e.equations[id] == nil {
		return nil, fmt.Errorf("no such id: %d", id)
	}
	return e.equations[id], nil
}

func (e *EquationMemory) GetEquationFromValue(value string) (*formula.Equation, error) {
	if e.equationFromValue[value] == nil {
		return nil, fmt.Errorf("no such value: %s", value)
	}
	return e.equationFromValue[value], nil
}

func (e *EquationMemory) GetExpression(id formula.ID) (*formula.Expression, error) {
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

func (e *EquationMemory) ListEquations() ([]*formula.Equation, error) {
	var equations []*formula.Equation
	for _, v := range e.equations {
		equations = append(equations, v)
	}
	return equations, nil
}

func (e *EquationMemory) ListExpressions() ([]*formula.Expression, error) {
	var expressions []*formula.Expression
	for _, v := range e.expressions {
		expressions = append(expressions, v)
	}
	return expressions, nil
}

func (e *EquationMemory) Upsert(equation *formula.Equation) (*formula.Equation, error) {
	if equation == nil {
		return nil, fmt.Errorf("expression is nil")
	}
	e.equations[equation.Id] = equation
	e.equationFromValue[equation.Value] = equation
	expression := e.parser.Parse(equation).(*formula.Expression)
	e.expressions[expression.EquationId] = expression
	return equation, nil
}

func (e *EquationMemory) InsertEquation(equation *formula.Equation) (*formula.Equation, error) {
	if equation == nil {
		return nil, fmt.Errorf("expression is nil")
	}
	e.equations[equation.Id] = equation
	e.equationFromValue[equation.Value] = equation
	return equation, nil
}

func (e *EquationMemory) InsertExpression(expression *formula.Expression) (*formula.Expression, error) {
	if expression == nil {
		return nil, fmt.Errorf("expression is nil")
	}
	e.expressions[expression.EquationId] = expression
	return expression, nil
}

func (e *EquationMemory) InsertConstants(id formula.ID, equationConstants *formula.EquationConstants) error {
	expression := e.expressions[id]
	if expression == nil {
		return fmt.Errorf("no such id: %d", id)
	}
	for _, constant := range equationConstants.Constants {
		expression.SetVariableConstant(constant)
	}

	return nil
}
