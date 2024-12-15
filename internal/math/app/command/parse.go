package command

import (
	"context"
	"fmt"
	"math-parser/internal/common/decorator"
	"math-parser/internal/math/domain/formula"
	"math-parser/internal/math/domain/formula/memory"
)

type Parse struct {
	Equation *formula.Equation
}
type ParseHandler decorator.CommandHandler[Parse]
type parseHandler struct {
	repo           formula.Repository
	equationMemory *memory.EquationMemory
	parser         formula.Parser
}

func NewParseHandler(repo formula.Repository, equationMemory *formula.EquationMemory, parser formula.Parser) decorator.CommandHandler[Parse] {
	if repo == nil {
		panic("nil repo")
	}
	if equationMemory == nil {
		panic("nil equationMemory")
	}
	return decorator.ApplyCommandDecorators[Parse](
		parseHandler{
			repo:           repo,
			equationMemory: equationMemory,
			parser:         parser,
		})
}

func (p parseHandler) Handle(ctx context.Context, cmd Parse) error {
	eq, err := p.repo.GetFromValue(ctx, cmd.Equation.Value)
	if eq == nil && err == nil {
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
	fmt.Printf("equation Id: %d, Constant: %s\n", eq.Id, eq.Value)

	expression := (p.parser.Parse(eq)).(*formula.Expression)
	print(expression.Description)
	p.equationMemory.InsertEquation(cmd.Equation)
	p.equationMemory.InsertExpression(expression)

	return nil
}
