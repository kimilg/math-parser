package command

import (
	"context"
	"fmt"
	"math-parser/internal/common/decorator"
	"math-parser/internal/math/domain/formula"
)

type Parse struct {
	Equation *formula.Equation
}
type ParseHandler decorator.CommandHandler[Parse]
type parseHandler struct {
	repo formula.Repository
}

func NewParseHandler(repo formula.Repository) decorator.CommandHandler[Parse] {
	if repo == nil {
		panic("nil repo")
	}
	return decorator.ApplyCommandDecorators[Parse](
		parseHandler{repo: repo})
}

func (p parseHandler) Handle(ctx context.Context, cmd Parse) error {
	eq, err := p.repo.GetFromValue(ctx, cmd.Equation.Value)
	if eq == nil && err != nil {
		eq, err = p.repo.Insert(ctx, cmd.Equation)
		if err != nil {
			return fmt.Errorf("internal server error: %w", err)
		}
	} else {
		if eq != nil {
			cmd.Equation.Id = eq.Id
			eq, err = p.repo.Update(ctx, cmd.Equation)
			if err != nil {
				return fmt.Errorf("internal server error: %w", err)
			}	
		}
	}
	fmt.Printf("equation Id: %d, Value: %s\n", eq.Id, eq.Value)

	equationExpression := formula.ParseEquation(eq)
	print(equationExpression.Description)
	
	return nil
}