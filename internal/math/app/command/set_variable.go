package command

import (
	"context"
	"fmt"
	"math-parser/internal/common/decorator"
	"math-parser/internal/math/domain/formula"
)

type SetVariable struct {
	VariableValue *formula.VariableValue
}

type SetVariableHandler decorator.CommandHandler[SetVariable]
type setVariableHandler struct {
	equationMemory *formula.EquationMemory
}

func NewSetVariable(equationMemory *formula.EquationMemory) SetVariableHandler {
	if equationMemory == nil {
		panic("nil equation memory")
	}
	return decorator.ApplyCommandDecorators[SetVariable](
		setVariableHandler{
			equationMemory: equationMemory,
		})
}

func (s setVariableHandler) Handle(ctx context.Context, cmd SetVariable) error {
	equation, err := s.equationMemory.GetEquationFromValue(cmd.VariableValue.Equation)
	if err != nil {
		return fmt.Errorf("fail to get equation from equationMemory: %w", err)
	}
	err = s.equationMemory.InsertVariableValue(equation.Id, cmd.VariableValue)
	if err != nil {
		return fmt.Errorf("fail to insert variable: %w", err)
	}

	return nil
}
