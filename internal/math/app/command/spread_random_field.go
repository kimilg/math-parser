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
	argumentMapper := make(map[argumentKey]field.IVector)

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
				argumentMapper[argumentKey{"displacement", "position"}] = displacementPosition

				for xi = 0; xi < field.Max; xi++ {
					for xj = 0; xj < field.Max; xj++ {
						for xk = 0; xk < field.Max; xk++ {
							forcePosition := field.Vector{xi, xj, xk, "force", "position"}
							argumentMapper[argumentKey{"force", "position"}] = forcePosition
							argumentMapper[argumentKey{"displacement", "time"}] = field.Scalar{1, "displacement", "time"}
							argumentMapper[argumentKey{"force", "time"}] = field.Scalar{0, "force", "time"}

							variableValues := []field.VariableValue{}
							for _, expression := range expressions {
								for _, element := range expression.Elements {
									switch e := element.(type) {
									case formula.Variable:
										variable := variables[e.Name]
										vectors := getArgumentVectors(argumentMapper, e)

										variableValues = append(variableValues, field.VariableValue{
											VariableDetail: e,
											Key:            buildKey(vectors),
											Value:          variable.ValueMapper[buildKey(vectors)],
										})
									case formula.Operator:
										// operator를 좌우 값을 기준으로 정할 수 없으니,, visitor를 또 돌려야 할까?
										// argument name을 알아야함. variable을 완벽히 알았다고 했을 때 visitor 돌리면 어떻게 대입할 수 있을지.
									}

								}
							}

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

func getArgumentVectors(argumentMapper map[argumentKey]field.IVector, variable formula.Variable) []field.IVector {
	vectors := []field.IVector{}
	for _, argument := range variable.Arguments {
		argVector := argumentMapper[argumentKey{argument.Category, argument.SubCategory}]
		vectors = append(vectors, argVector)
	}
	return vectors
}

type argumentKey struct {
	Category    string
	SubCategory string
}
