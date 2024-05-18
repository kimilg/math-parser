package adapters

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"math-parser/db"
	"math-parser/internal/math/domain/formula"
)

type Repository struct {
	queries *db.Queries
}

func NewRepository(queries *db.Queries) *Repository {
	return &Repository{
		queries: queries,
	}
}

func (r *Repository) Get(ctx context.Context, id formula.ID) (*formula.Equation, error) {
	eqn, err := r.queries.GetEquation(ctx, int64(id))
	if err != nil {
		return nil, fmt.Errorf("error reading from database: %w", err)
	}
	var variables []*formula.Variable
	err = json.Unmarshal(eqn.Variables, &variables)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling variables: %w", err)
	}
	
	return &formula.Equation{
		Id:    formula.ID(eqn.ID),
		Value: eqn.Value,
		Category: eqn.Category,
		Variables: variables,
		Cause: pgtypeToString(eqn.Cause),
		Effect: pgtypeToString(eqn.Effect),
	}, nil
}

func (r *Repository) GetFromValue(ctx context.Context, value string) (*formula.Equation, error) {
	eqn, err := r.queries.GetEquationFromValue(ctx, value)
	if err != nil {
		return nil, fmt.Errorf("error reading from database: %w", err)
	}
	var variables []*formula.Variable
	err = json.Unmarshal(eqn.Variables, &variables)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling variables: %w", err)
	}
	
	return &formula.Equation{
		Id:    formula.ID(eqn.ID),
		Value: eqn.Value,
		Category: eqn.Category,
		Variables: variables,
		Cause: pgtypeToString(eqn.Cause),
		Effect: pgtypeToString(eqn.Effect),
	}, nil
}

func (r *Repository) List(ctx context.Context) ([]*formula.Equation, error) {
	eqns, err := r.queries.ListEquations(ctx)
	if err != nil {
		return nil, fmt.Errorf("error reading from database: %w", err)
	}
	var equations []*formula.Equation
	for _, eqn := range eqns {
		var variables []*formula.Variable
		err = json.Unmarshal(eqn.Variables, &variables)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling variables: %w", err)
		}
		
		equations = append(equations,
			&formula.Equation{
				Id:    formula.ID(eqn.ID),
				Value: eqn.Value,
				Category: eqn.Category,
				Variables: variables,
				Cause: pgtypeToString(eqn.Cause),
				Effect: pgtypeToString(eqn.Effect),
			})
	}
	return equations, nil
}

func (r *Repository) Insert(ctx context.Context, equation *formula.Equation) (*formula.Equation, error) {
	eqn, err := r.queries.InsertEquation(ctx, 
		db.InsertEquationParams{
			Value: equation.Value, 
			Category: equation.Category, 
			Cause: pgtype.Text{String: equation.Cause, Valid: true},
			Effect: pgtype.Text{String: equation.Effect, Valid: true},
		})
	if err != nil {
		return nil, fmt.Errorf("error creating equation: %w", err)
	}
	
	for _, variable := range equation.Variables {
		data, err := json.Marshal(variable.Arguments)
		if err != nil {
			return nil, fmt.Errorf("error marshaling variable.Arguments: %w", err)
		}
		_, err = r.queries.InsertVariable(ctx, 
			db.InsertVariableParams{
				Name: variable.Name,
				Vcategory: variable.Vcategory,
				Arguments: data,
				EquationID: eqn.ID,
			})
		if err != nil {
			return nil, fmt.Errorf("error inserting equation_variable: %w", err)
		}
	}

	return equation, nil
}

func (r *Repository) Update(ctx context.Context, equation *formula.Equation) (*formula.Equation, error) {
	eqn, err := r.queries.UpdateEquation(ctx, 
		db.UpdateEquationParams{
		ID: int64(equation.Id),
		Value: equation.Value,
		Category: equation.Category,
		Cause: pgtype.Text{String: equation.Cause, Valid: true},
		Effect: pgtype.Text{String: equation.Effect, Valid: true},
		})
	if err != nil {
		return nil, fmt.Errorf("error updating equation: %w", err)
	}

	for _, variable := range equation.Variables {
		data, err := json.Marshal(variable.Arguments)
		if err != nil {
			return nil, fmt.Errorf("error marshaling variable.Arguments: %w", err)
		}
		_, err = r.queries.InsertVariable(ctx,
			db.InsertVariableParams{
				Name: variable.Name,
				Vcategory: variable.Vcategory,
				Arguments: data,
				EquationID: eqn.ID,
			})
		if err != nil {
			return nil, fmt.Errorf("error inserting equation_variable: %w", err)
		}
	}
	
	return equation, nil
}

func (r *Repository) Delete(ctx context.Context, id formula.ID) error {
	err := r.queries.DeleteEquation(ctx, int64(id))
	if err != nil {
		return fmt.Errorf("error deleting equation: %w", err)
	}
	return nil
}

func pgtypeToString(text pgtype.Text) string {
	if text.Valid {
		return text.String
	}
	return ""
}

