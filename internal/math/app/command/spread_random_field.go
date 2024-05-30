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
	repo           formula.Repository
	equationMemory *formula.EquationMemory
}

func NewSpreadRandomFieldHandler(repo formula.Repository, equationMemory *formula.EquationMemory) decorator.CommandHandler[SpreadRandomField] {
	if repo == nil {
		panic("repo nil")
	}
	if equationMemory == nil {
		panic("equationMemory nil")
	}
	return decorator.ApplyCommandDecorators[SpreadRandomField](
		spreadRandomFieldHandler{
			repo:           repo,
			equationMemory: equationMemory,
		})
}

func (s spreadRandomFieldHandler) Handle(ctx context.Context, cmd SpreadRandomField) error {
	variables := make(map[string]*field.Variable)

	equations, err := s.equationMemory.ListEquations()
	if err != nil {
		return fmt.Errorf("fail to get equations from equationMemory: %w", err)
	}
	for _, equation := range equations {
		for _, v := range equation.Variables {
			variables[v.Name] = field.NewVariable(v.Name, v.Vcategory)
		}
	}

	expressions, err := s.equationMemory.ListExpressions()
	if err != nil {
		return fmt.Errorf("fail to get expressions from equationMemory: %w", err)
	}
	var targetExpr []*formula.Expression

	for _, expression := range expressions {
		if expression.Category == "field_making_rule" {
			targetExpr = append(targetExpr, expression)
		}
	}

	var i, j, k, xi, xj, xk field.Pos
	for i = 0; i < field.Max; i++ {
		for j = 0; j < field.Max; j++ {
			for k = 0; k < field.Max; k++ {
				displacementPosition := field.Vector{i, j, k, "displacement", "position"}

				for xi = 0; xi < field.Max; xi++ {
					for xj = 0; xj < field.Max; xj++ {
						for xk = 0; xk < field.Max; xk++ {
							forcePosition := field.Vector{xi, xj, xk, "force", "position"}

							key := buildKey([]field.IVector{
								displacementPosition,
								field.Scalar{1, "displacement", "time"},
								forcePosition,
								field.Scalar{0, "force", "time"},
							})

							for _, expression := range expressions {
								for _, element := range expression.Elements {
									switch e := element.(type) {
									case formula.Variable:
										variable := variables[e.Name]
										for _, argument := range e.Arguments {
											// sum을 만들어서,, 거기에 expr를 실제 연산으로 바꾼다.
											// 여기서 이 variable의 값이 존재하는지 확인 가능.
											// 미지수가 하나만 남는다면 그 값을 지정할 수 있을 것이다.
											// 미지수가 두개라면 random 값을 대입해야 할 듯.
											// 첫 번째 argument, 
											// 두 번째 argument,
											// ...
											argument.Category, argument.SubCategory

										}
									}
								}
							}

							print(FieldMap[key])

						}
					}
				}
			}
		}
	}

	return nil
}

func buildKey(iVectors []field.IVector) string {
	return fmt.Sprintf("%v", iVectors)
}
