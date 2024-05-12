package adapters

import (
	"context"
	"fmt"
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
	return &formula.Equation{
		Id:    formula.ID(eqn.ID), 
		Value: eqn.Value,
	}, nil
}

func (r *Repository) GetFromValue(ctx context.Context, value string) (*formula.Equation, error) {
	eqn, err := r.queries.GetEquationFromValue(ctx, value)
	if err != nil {
		return nil, fmt.Errorf("error reading from database: %w", err)
	}
	return &formula.Equation{
		Id:    formula.ID(eqn.ID),
		Value: eqn.Value,
	}, nil
}

func (r *Repository) List(ctx context.Context) ([]*formula.Equation, error) {
	eqns, err := r.queries.ListEquations(ctx); 
	if err != nil {
		return nil, fmt.Errorf("error reading from database: %w", err)
	}
	var equations []*formula.Equation
	for _, eqn := range eqns {
		equations = append(equations, 
			&formula.Equation{
			Id:    formula.ID(eqn.ID), 
			Value: eqn.Value,
			})
	}
	return equations, nil
}

func (r *Repository) Create(ctx context.Context, value string, category string) (*formula.Equation, error) {
	eqn, err := r.queries.CreateEquation(ctx, db.CreateEquationParams{Value: value, Category: category})
	if err != nil {
		return nil, fmt.Errorf("error creating equation: %w", err)
	}
	return &formula.Equation{
		Id:    formula.ID(eqn.ID),
		Value: eqn.Value,
	}, nil
}

func (r *Repository) Update(ctx context.Context, id formula.ID, value string) (*formula.Equation, error) {
	eqn, err := r.queries.UpdateEquation(ctx, db.UpdateEquationParams{int64(id), value})
	if err != nil {
		return nil, fmt.Errorf("error updating equation: %w", err)
	}
	return &formula.Equation{
		Id:    formula.ID(eqn.ID),
		Value: eqn.Value,
	}, nil
}

func (r *Repository) Delete(ctx context.Context, id formula.ID) error {
	err := r.queries.DeleteEquation(ctx, int64(id))
	if err != nil {
		return fmt.Errorf("error deleting equation: %w", err)
	}
	return nil
}

