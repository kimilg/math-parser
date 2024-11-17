package calculator

import (
	"fmt"
	"math-parser/internal/math/app/command/entity"
	"math-parser/internal/math/domain/field"
	"math-parser/internal/math/domain/formula"
	"strconv"
)

func DynamicLoop(mathInfo *entity.MathInfo) error {
	equation := mathInfo.Equation
	expression := mathInfo.Expression
	variableMapper := mathInfo.VariableMapper
	argumentKinds := mathInfo.ArgumentKinds
	argumentMapper := mathInfo.ArgumentMapper

	argumentKind := formula.PopArgumentKind(argumentKinds)
	if argumentKind == nil {
		return assignValue(mathInfo)
	}
	dimension := field.Dimension[argumentKind.SubCategory]
	if dimension == 0 {
		return fmt.Errorf("error while generating loop: there is no loop size definition for SubCategory" +
			argumentKind.SubCategory)
	}
	switch dimension {
	case 3:
		return dynamicLoopThree(equation, expression, variableMapper, argumentKinds, argumentKind, argumentMapper)
	case -1:
		return dynamicLoopNone(equation, expression, variableMapper, argumentKinds, argumentKind, argumentMapper)
	default:
		return fmt.Errorf("error while generating loop: there is no loop size other than 1 and 3")
	}
}

func assignValue(mathInfo *entity.MathInfo) error {
	argumentMapper := mathInfo.ArgumentMapper
	variableMapper := mathInfo.VariableMapper
	expression := mathInfo.Expression

	result := traverse(expression, argumentMapper, variableMapper)
}

func traverse(expression *formula.Expression, argumentMapper map[string]field.IVector, variableValueMapper map[string]*field.Variable) *formula.Result {
	if len(expression.Elements) == 3 {
		if operator, ok := expression.Elements[1].(byte); ok {
			left := traverse(expression.Elements[0].(*formula.Expression), argumentMapper, variableValueMapper)
			right := traverse(expression.Elements[2].(*formula.Expression), argumentMapper, variableValueMapper)
			switch operator {
			case formula.EQUAL:
				return left.Equal(right)
			case formula.PLUS:
				return left.Plus(right)
			case formula.MINUS:
				return left.Minus(right)
			case formula.MULT:
				return left.Mult(right)

			default:
				panic("unhandled default case")
			}
		}
	}

	for _, element := range expression.Elements {
		switch e := element.(type) {
		case formula.Variable:
			variableValue := variableValueMapper[e.Name]
			key, err := BuildKey(argumentMapper, e)
			if err != nil {
				return &formula.Result{Unknowns: []*formula.Unknown{{Name: e.Name, Value: nil}}}
			}

			valueVector := variableValue.Mapper[key]
			return &formula.Result{Unknowns: []*formula.Unknown{{Name: e.Name, Value: valueVector}}}

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
				err := DynamicLoop(equation, expression, variables, copySlice(argumentKinds), arguments)
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
	return DynamicLoop(equation, expression, variables, argumentKinds, arguments)
}

func BuildKey(argumentMapper map[string]field.IVector, variable formula.Variable) (string, error) {
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

func GetForce() []field.IVector {
	vectors := []*field.Vector{{1, 0, 0, "force", "amount"},
		{0, 1, 0, "force", "amount"},
		{0, 0, 1, "force", "amount"}}
	iVectors := make([]field.IVector, len(vectors))
	for i, vector := range vectors {
		iVectors[i] = field.IVector(vector)
	}
	return iVectors
}
