package command

import (
	"context"
	"fmt"
	"math-parser/internal/common/decorator"
	"math-parser/internal/math/domain/formula"
)

type Parse struct {
	Equation string
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
	eq, err := p.repo.GetFromValue(ctx, cmd.Equation)
	if eq == nil && err != nil {
		eq, err = p.repo.Create(ctx, cmd.Equation)
		if err != nil {
			return fmt.Errorf("internal server error")
		}
	}
	fmt.Printf("equation Id: %d, Value: %s\n", eq.Id, eq.Value)

	equation := formula.ParseEquation(eq.Id, eq.Value)
	print(equation.Description)
	
	return nil
}