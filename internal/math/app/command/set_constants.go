package command

import (
	"context"
	"fmt"
	"math-parser/internal/common/decorator"
	"math-parser/internal/math/domain/formula"
	"math-parser/internal/math/domain/formula/memory"
)

type SetConstant struct {
	EquationConstants *formula.EquationConstants
}

type SetConstantsHandler decorator.CommandHandler[SetConstant]
type setConstantsHandler struct {
	equationMemory *memory.EquationMemory
}

func NewSetConstants(equationMemory *memory.EquationMemory) SetConstantsHandler {
	if equationMemory == nil {
		panic("nil equation memory")
	}
	return decorator.ApplyCommandDecorators[SetConstant](
		setConstantsHandler{
			equationMemory: equationMemory,
		})
}

func (s setConstantsHandler) Handle(ctx context.Context, cmd SetConstant) error {
	equation, err := s.equationMemory.GetEquationFromValue(cmd.EquationConstants.Equation)
	if err != nil {
		return fmt.Errorf("fail to get equation from equationMemory: %w", err)
	}

	// 이 사이에 constant가 외부 adapter 용에서 또 다른 entity로 바뀌어야할 것 같고, 그리고 서비스 레이어에서 비지니스용 constant로 바꾸는게 필요.

	err = s.equationMemory.InsertConstants(equation.Id, cmd.EquationConstants)
	if err != nil {
		return fmt.Errorf("fail to insert variable: %w", err)
	}

	return nil
}
