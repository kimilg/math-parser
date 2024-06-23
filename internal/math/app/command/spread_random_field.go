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
			variableValueMapper := make(map[string]*field.VariableValue)
			for _, v := range equation.Variables {
				variableValueMapper[v.Name] = field.NewVariableValue(v.Name, v.Vcategory)
			}
			err := s.calculateByField(equation.Value, variableValueMapper, expression)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func assignValue(argumentMapper map[string]field.IVector, eqStr string, 
	variableValueMapper map[string]*field.VariableValue, expression *formula.Expression) error {
	
	result := traverse(expression, argumentMapper, variableValueMapper)
}

func traverse(expression *formula.Expression, argumentMapper map[string]field.IVector, 
	variableValueMapper map[string]*field.VariableValue) *formula.Result {
	
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

func dynamicLoop(argumentConcretes []*formula.ArgumentConcrete, argumentMapper map[string]field.IVector, eqStr string, 
	variableValueMapper map[string]*field.VariableValue, expression *formula.Expression) error {
	
	argumentConcrete := formula.PopArgumentConcrete(argumentConcretes)
	if argumentConcrete == nil {
		return assignValue(argumentMapper, eqStr, variableValueMapper, expression)
	}
	dim := field.Dimension[argumentConcrete.SubCategory]
	if dim == 0 {
		return fmt.Errorf("error while generating loop: there is no loop size definition for SubCategory" + 
			argumentConcrete.SubCategory)
	}
	switch dim {
	case 3:
		return dynamicLoopThree(argumentConcretes, argumentConcrete, argumentMapper, eqStr, variableValueMapper, expression)
	case -1:
		return dynamicLoopNone(argumentConcretes, argumentConcrete, argumentMapper, eqStr, variableValueMapper, expression)
	default:
		return fmt.Errorf("error while generating loop: there is no loop size other than 1 and 3")
	}
}

func dynamicLoopThree(argumentConcretes []*formula.ArgumentConcrete, argumentConcrete *formula.ArgumentConcrete, 
	argumentMapper map[string]field.IVector, eqStr string, 
	variableValueMapper map[string]*field.VariableValue, expression *formula.Expression) error {
	
	var i, j, k field.Pos
	for i = 0; i < field.Max; i++ {
		for j = 0; j < field.Max; j++ {
			for k = 0; k < field.Max; k++ {
				argumentMapper[argumentConcrete.Name] = &field.Vector{i, j, k, "", argumentConcrete.SubCategory}
				err := dynamicLoop(copySlice(argumentConcretes), argumentMapper, eqStr, variableValueMapper, expression)
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

func dynamicLoopNone(argumentConcretes []*formula.ArgumentConcrete, argumentConcrete *formula.ArgumentConcrete,
	argumentMapper map[string]field.IVector, eqStr string,
	variableValueMapper map[string]*field.VariableValue, expression *formula.Expression) error {
	
	val, err := strconv.Atoi(argumentConcrete.Name)
	if err != nil {
		return fmt.Errorf("error while converting %s to int: %w", argumentConcrete.Name, err)
	}
	argumentMapper[argumentConcrete.Name] = &field.Scalar{X: field.Val(val), SubCategory: argumentConcrete.SubCategory}
	return dynamicLoop(argumentConcretes, argumentMapper, eqStr, variableValueMapper, expression)
}


func (s spreadRandomFieldHandler) calculateByField(eqStr string, variableValueMapper map[string]*field.VariableValue, 
	expression *formula.Expression) error {

	argumentMapper := make(map[string]field.IVector)
	err := dynamicLoop(expression.GetArgumentConcretes(), argumentMapper, eqStr, variableValueMapper, expression)
	if err != nil {
		return fmt.Errorf("error while calculating field %s: %w", eqStr, err)
	}
	
	
	var i, j, k, xi, xj, xk field.Pos
	var dir int
	for i = 0; i < field.Max; i++ {
		for j = 0; j < field.Max; j++ {
			for k = 0; k < field.Max; k++ {
				displacementPosition := field.Vector{i, j, k, "displacement", "position"}
				argumentMapper[formula.ArgumentKey{"displacement", "position"}] = &displacementPosition

				for xi = 0; xi < field.Max; xi++ {
					for xj = 0; xj < field.Max; xj++ {
						for xk = 0; xk < field.Max; xk++ {
							for dir = 0; dir < 3; dir++ {

								argumentMapper[formula.ArgumentKey{"force", "amount"}] = getForce()[dir]
								forcePosition := field.Vector{xi, xj, xk, "force", "position"}
								argumentMapper[formula.ArgumentKey{"force", "position"}] = &forcePosition
								argumentMapper[formula.ArgumentKey{"displacement", "time"}] = &field.Scalar{1, "displacement", "time"}
								argumentMapper[formula.ArgumentKey{"force", "time"}] = &field.Scalar{0, "force", "time"}

								
								
								
								assignParser := s.parser.GetParser(eqStr)
								if assignVisitor, ok := s.visitor.(*visitor.AssignVisitorImpl); ok {
									assignVisitor.ArgumentMapper = argumentMapper
									assignVisitor.VariableValueMapper = variableValueMapper
									assignVisitor.Visit(assignParser.Equation())
								}

								for _, element := range expression.Elements {
									switch e := element.(type) {
									case formula.Variable:
										variableValue := variableValueMapper[e.Name]
										value := variableValue.Mapper[buildKey(argumentMapper, e)]
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