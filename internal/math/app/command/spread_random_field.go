package command

import (
	"context"
	"math-parser/internal/common/decorator"
	"math-parser/internal/math/domain/field"
	"math-parser/internal/math/domain/formula"
)

type SpreadRandomField struct {
	
}

type SpreadRandomFieldHandler decorator.CommandHandler[SpreadRandomField]
type spreadRandomFieldHandler struct {
	repo formula.Repository
}

func NewSpreadRandomFieldHandler(repo formula.Repository) decorator.CommandHandler[SpreadRandomField] {
	return decorator.ApplyCommandDecorators[SpreadRandomField](spreadRandomFieldHandler{})
}

func (s spreadRandomFieldHandler) Handle(ctx context.Context, cmd SpreadRandomField) error {
	fieldData := field.NewField()
	
	var xi, xj, xk, i, j, k field.Pos
	for xi = 0; xi < field.Max; xi++ {
		for xj = 0; xj < field.Max; xj++ {
			for xk = 0; xk < field.Max; xk++ {
				forcePosition := field.Position{xi, xj, xk}
				
					
				for i = 0; i < field.Max; i++ {
					for j = 0; j < field.Max; j++ {
						for k = 0; k < field.Max; k++ {
							
						}
					}
				}
			}
		}
	}
	
	return nil
}