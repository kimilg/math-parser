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
	parser         formula.Parser
}

func NewSpreadRandomFieldHandler(repo formula.Repository, equationMemory *formula.EquationMemory, parser formula.Parser) decorator.CommandHandler[SpreadRandomField] {
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
			parser: parser,
		})
}

func (s spreadRandomFieldHandler) Handle(ctx context.Context, cmd SpreadRandomField) error {
	pairs, err := s.equationMemory.List()
	if err != nil {
		return fmt.Errorf("fail to get equation and expression pair from equationMemory: %w", err)
	}
	for _, pair := range pairs {
		equation := pair.Equation
		expression := pair.Expression
		if equation.Category == "field_making_rule" && expression.Category == "field_making_rule" {
			variableValueMapper := make(map[string]*field.VariableValue)
			for _, v := range equation.Variables {
				variableValueMapper[v.Name] = field.NewVariableValue(v.Name, v.Vcategory)
			}
			calculateByField(variableValueMapper, expression)
		}
	}
	return nil
}

func calculateByField(variableValueMapper map[string]*field.VariableValue, expression *formula.Expression) {
	argumentMapper := make(map[formula.ArgumentKey]field.IVector)
	var i, j, k, xi, xj, xk field.Pos
	for i = 0; i < field.Max; i++ {
		for j = 0; j < field.Max; j++ {
			for k = 0; k < field.Max; k++ {
				displacementPosition := field.Vector{i, j, k, "displacement", "position"}
				argumentMapper[formula.ArgumentKey{"displacement", "position"}] = displacementPosition

				for xi = 0; xi < field.Max; xi++ {
					for xj = 0; xj < field.Max; xj++ {
						for xk = 0; xk < field.Max; xk++ {
							forcePosition := field.Vector{xi, xj, xk, "force", "position"}
							argumentMapper[formula.ArgumentKey{"force", "position"}] = forcePosition
							argumentMapper[formula.ArgumentKey{"displacement", "time"}] = field.Scalar{1, "displacement", "time"}
							argumentMapper[formula.ArgumentKey{"force", "time"}] = field.Scalar{0, "force", "time"}
							
							
							
							
							for _, element := range expression.Elements {
								switch e := element.(type) {
								case formula.Variable:
									variableValue := variableValueMapper[e.Name]
									vectors := getArgumentVectors(argumentMapper, e)
									
									value := variableValue.Mapper[buildKey(vectors)]
									print(value)
								}
							}
							// argument 이름도 포함되어야 parsing 할 때 적용될 수 있을 듯. 

						}
					}
				}
			}
		}
	}
}

func buildKey(iVectors []field.IVector) string {
	return fmt.Sprintf("%v", iVectors)
}

func getArgumentVectors(argumentMapper map[formula.ArgumentKey]field.IVector, variable formula.Variable) []field.IVector {
	vectors := []field.IVector{}
	for _, argument := range variable.Arguments {
		argVector := argumentMapper[formula.ArgumentKey{argument.Category, argument.SubCategory}]
		vectors = append(vectors, argVector)
	}
	return vectors
}
