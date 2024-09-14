package command

import (
	"context"
	"fmt"
	"math-parser/internal/common/decorator"
	"math-parser/internal/math/domain/field"
	"math-parser/internal/math/domain/formula"
	"math-parser/internal/math/domain/visitor"
	"math-parser/parser"
	"strconv"
)

type SpreadRandomField struct {
}

type SpreadRandomFieldHandler decorator.CommandHandler[SpreadRandomField]
type spreadRandomFieldHandler struct {
	repo           formula.Repository
	equationMemory *formula.EquationMemory
	parser         formula.Parser
	visitor		   parser.FormulaVisitor
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
			parser: parser,
			visitor: visitor,
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
			variables := make(map[string]*field.Variable)
			for _, v := range equation.Variables {
				variables[v.Name] = field.NewVariable(v.Name, v.Vcategory)
			}
			err := s.calculateByField(equation.Value, expression, variables)
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
	err := dynamicLoop(equation, expression, variables, expression.GetArgumentKinds(), arguments)
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

								arguments[formula.ArgumentKey{"force", "amount"}] = getForce()[dir]
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
										value := variableValue.Mapper[buildKey(arguments, e)]
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

func dynamicLoop(equation string, expression *formula.Expression, variables map[string]*field.Variable, 
	argumentKinds []*formula.ArgumentKind, arguments map[string]field.IVector) error {

	argumentKind := formula.PopArgumentKind(argumentKinds)
	if argumentKind == nil {
		return assignValue(arguments, equation, variables, expression)
	}
	dimension := field.Dimension[argumentKind.SubCategory]
	if dimension == 0 {
		return fmt.Errorf("error while generating loop: there is no loop size definition for SubCategory" +
			argumentKind.SubCategory)
	}
	switch dimension {
	case 3:
		return dynamicLoopThree(equation, expression, variables, argumentKinds, argumentKind, arguments)
	case -1:
		return dynamicLoopNone(equation, expression, variables, argumentKinds, argumentKind, arguments)
	default:
		return fmt.Errorf("error while generating loop: there is no loop size other than 1 and 3")
	}
}

func assignValue(argumentMapper map[string]field.IVector, eqStr string, 
	variableValueMapper map[string]*field.Variable, expression *formula.Expression) error {
	
	result := traverse(expression, argumentMapper, variableValueMapper)
}

func traverse(expression *formula.Expression, argumentMapper map[string]field.IVector, 
	variableValueMapper map[string]*field.Variable) *formula.Result {
	
	if len(expression.Elements) == 3 {
		if operator, ok := expression.Elements[1].(byte); ok {
			left := traverse(expression.Elements[0].(*formula.Expression), argumentMapper, variableValueMapper)
			right := traverse(expression.Elements[2].(*formula.Expression), argumentMapper, variableValueMapper)
			switch operator {
			case formula.PLUS:
				return left.Plus(right)
			case formula.MINUS:
				return left.Minus(right)
			case formula.MULT:
				
			default:
				panic("unhandled default case")
				
			}
		}
	}
	
	for _, element := range expression.Elements {
		switch e := element.(type) {
		case formula.Variable:
			variableValue := variableValueMapper[e.Name]
			key, err := buildKey(argumentMapper, e); 
			if err != nil {
				return &formula.Result{Unknowns: []*formula.Unknown{{Name: e.Name, Coefficient: 1, Value: nil}}}
			}
			
			valueVector := variableValue.Mapper[key]
			return &formula.Result{Unknowns: []*formula.Unknown{{Name: e.Name, Coefficient: 1, Value: valueVector}}}
			
		case formula.Constant:
			return &formula.Result{Constant: e.Value}
		
			
		}
	}
}

func dynamicLoopThree(equation string, expression *formula.Expression, variables map[string]*field.Variable,
	argumentKinds []*formula.ArgumentKind, argumentKind *formula.ArgumentKind, arguments map[string]field.IVector) error {
	
	var i, j, k field.Pos
	for i = 0; i < field.Max; i++ {
		for j = 0; j < field.Max; j++ {
			for k = 0; k < field.Max; k++ {
				arguments[argumentKind.Name] = &field.Vector{i, j, k, "", argumentKind.SubCategory}
				err := dynamicLoop(equation, expression, variables, copySlice(argumentKinds), arguments)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func copySlice[T any](orig []T) []T {
	return append([]T(nil), orig...)
}

func dynamicLoopNone(equation string, expression *formula.Expression, variables map[string]*field.Variable, 
	argumentKinds []*formula.ArgumentKind, argumentKind *formula.ArgumentKind, arguments map[string]field.IVector) error {
	
	val, err := strconv.Atoi(argumentKind.Name)
	if err != nil {
		return fmt.Errorf("error while converting %s to int: %w", argumentKind.Name, err)
	}
	arguments[argumentKind.Name] = &field.Scalar{X: field.Val(val), SubCategory: argumentKind.SubCategory}
	return dynamicLoop(equation, expression, variables, argumentKinds, arguments)
}

func buildKey(argumentMapper map[string]field.IVector, variable formula.Variable) (string, error) {
	vectors := []field.IVector{}
	for _, argument := range variable.Arguments {
		argVector := argumentMapper[argument.Name]
		if argVector == nil {
			return "", fmt.Errorf("cannot find argument from argumentMapper")
		}
		vectors = append(vectors, argVector)
	}
	return fmt.Sprintf("%v", vectors), nil
}

func getForce() []field.IVector {
	vectors := []*field.Vector{{1,0,0,"force","amount"},
		{0,1,0,"force","amount"},
		{0,0,1,"force","amount"}}
	iVectors := make([]field.IVector, len(vectors))
	for i, vector := range vectors {
		iVectors[i] = field.IVector(vector)
	}
	return iVectors
}