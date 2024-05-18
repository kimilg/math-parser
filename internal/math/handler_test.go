package math

import (
	"context"
	"encoding/json"
	"math-parser/internal/math/domain/formula"
	"testing"
)

func TestNew(t *testing.T) {
	var variables []*formula.Variable
	jsonStr := `[
        {
            "name": "G",
            "vcategory": "displacement",
            "arguments": [
                {
                    "seq": 1,
                    "category": "position"
                },
                {
                    "seq": 2,
                    "category": "time"
                },
                {
                    "seq": 3,
                    "category": "position",
                    "subcategory": "force"
                },
                {
                    "seq": 4,
                    "category": "time",
                    "subcategory": "force"
                }
            ]
        }
    ]`
	
	err := json.Unmarshal([]byte(jsonStr), &variables)
	if err != nil {}
	print(variables)
}


type mockEquationRepository struct {}

func (r *mockEquationRepository) Create(ctx context.Context, value string) (*formula.Equation, error) {
	return &formula.Equation{
		Id: 0,
		Value: value,
	}, nil
}
func (r *mockEquationRepository) Get(ctx context.Context, id formula.ID) (*formula.Equation, error) {
	return nil, nil
}
func (r *mockEquationRepository) GetFromValue(ctx context.Context, value string) (*formula.Equation, error) {
	return nil, nil
}
func (r *mockEquationRepository) List(ctx context.Context) ([]*formula.Equation, error) {
	return nil, nil
}
func (r *mockEquationRepository) Update(ctx context.Context, id formula.ID, value string) (*formula.Equation, error) {
	return nil, nil
}
func (r *mockEquationRepository) Delete(ctx context.Context, id formula.ID) error {
	return nil
}

func TestParse(t *testing.T) {
	
}