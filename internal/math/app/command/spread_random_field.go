package command

import (
	"context"
	"fmt"
	"math-parser/internal/common/decorator"
	"math-parser/internal/math/app/command/calculator"
	"math-parser/internal/math/app/command/entity"
	"math-parser/internal/math/domain/field"
	"math-parser/internal/math/domain/formula"
	"math-parser/internal/math/domain/visitor"
	"math-parser/parser"
)

type SpreadRandomField struct {
}

type SpreadRandomFieldHandler decorator.CommandHandler[SpreadRandomField]
type spreadRandomFieldHandler struct {
	repo           formula.Repository
	equationMemory *formula.EquationMemory
	parser         formula.Parser
	visitor        parser.FormulaVisitor
}

func NewSpreadRandomFieldHandler(repo formula.Repository, equationMemory *formula.EquationMemory, parser formula.Parser,
	visitor parser.FormulaVisitor) decorator.CommandHandler[SpreadRandomField] {
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
			parser:         parser,
			visitor:        visitor,
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
			//여기서 variable mapper를 또 다른 controller를 통해 가져와야 할 듯. 이미 메모리에 담고 있어야 할듯. 하다.
			// variableMapper가 필요없을 수 있다. 이제 equation 안의 variable에 값이 들어가 있으니. 그냥 그걸 쓰면 되겠다.
			variableMapper := make(map[string]*field.Variable)
			for _, v := range equation.Variables {
				variableMapper[v.Name] = field.NewVariable(v.Name, v.Vcategory)
			}
			err := s.calculateByField(equation.Value, expression, variableMapper)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s spreadRandomFieldHandler) calculateByField(equation string, expression *formula.Expression,
	variables map[string]*field.Variable) error {

	arguments := make(map[string]field.IVector)
	err := calculator.DynamicLoop(&entity.MathInfo{equation, expression, variables, expression.GetArgumentKinds(), arguments})
	if err != nil {
		return fmt.Errorf("error while calculating field %s: %w", equation, err)
	}

	var i, j, k, xi, xj, xk field.Pos
	var dir int
	for i = 0; i < field.Max; i++ {
		for j = 0; j < field.Max; j++ {
			for k = 0; k < field.Max; k++ {
				displacementPosition := field.Vector{i, j, k, "displacement", "position"}
				arguments[formula.ArgumentKey{"displacement", "position"}] = &displacementPosition

				for xi = 0; xi < field.Max; xi++ {
					for xj = 0; xj < field.Max; xj++ {
						for xk = 0; xk < field.Max; xk++ {
							for dir = 0; dir < 3; dir++ {

								arguments[formula.ArgumentKey{"force", "amount"}] = calculator.GetForce()[dir]
								forcePosition := field.Vector{xi, xj, xk, "force", "position"}
								arguments[formula.ArgumentKey{"force", "position"}] = &forcePosition
								arguments[formula.ArgumentKey{"displacement", "time"}] = &field.Scalar{1, "displacement", "time"}
								arguments[formula.ArgumentKey{"force", "time"}] = &field.Scalar{0, "force", "time"}

								assignParser := s.parser.GetParser(equation)
								if assignVisitor, ok := s.visitor.(*visitor.AssignVisitorImpl); ok {
									assignVisitor.ArgumentMapper = arguments
									assignVisitor.VariableValueMapper = variables
									assignVisitor.Visit(assignParser.Equation())
								}

								for _, element := range expression.Elements {
									switch e := element.(type) {
									case formula.Variable:
										variableValue := variables[e.Name]
										value := variableValue.Mapper[calculator.BuildKey(arguments, e)]
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
}
