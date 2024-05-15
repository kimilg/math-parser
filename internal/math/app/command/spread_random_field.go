package command

import (
	"context"
	"fmt"
	"math-parser/internal/common/decorator"
	"math-parser/internal/math/domain/field"
	"math-parser/internal/math/domain/formula"
)

type SpreadRandomField struct {
	
}

type SpreadRandomFieldHandler decorator.CommandHandler[SpreadRandomField]
type spreadRandomFieldHandler struct {
	repo formula.Repository
	equationMemory formula.EquationMemory
}

func NewSpreadRandomFieldHandler(repo formula.Repository, equationMemory formula.EquationMemory) decorator.CommandHandler[SpreadRandomField] {
	return decorator.ApplyCommandDecorators[SpreadRandomField](
		spreadRandomFieldHandler{
			repo: repo,
			equationMemory: equationMemory,
		})
}

func (s spreadRandomFieldHandler) Handle(ctx context.Context, cmd SpreadRandomField) error {
	Fields := field.NewFields()
	expressions, err := s.equationMemory.List()
	if err != nil {
		return fmt.Errorf("fail to get equationMemory: %w", err)
	}
	var targetExpr []*formula.Expression
	
	for _, expression := range expressions {
		if expression.Category == "field_making_rule" {
			targetExpr = append(targetExpr, expression)
		}
	}
	
	var xi, xj, xk, i, j, k field.Pos
	for xi = 0; xi < field.Max; xi++ {
		for xj = 0; xj < field.Max; xj++ {
			for xk = 0; xk < field.Max; xk++ {
				forcePosition := field.Position{xi, xj, xk}
				
					
				for i = 0; i < field.Max; i++ {
					for j = 0; j < field.Max; j++ {
						for k = 0; k < field.Max; k++ {
							displacementPosition := field.Position{i, j, k}
							
							print(Fields.FieldMap[field.DFPosition{displacementPosition, forcePosition}].Force.Y)
							
						}
					}
				}
			}
		}
	}
	
	return nil
}