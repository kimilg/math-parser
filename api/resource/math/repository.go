package math

import (
	"context"
	"fmt"
	"math-parser/db"
)

type Repository struct {
	queries *db.Queries
}

func NewRepository(queries *db.Queries) *Repository {
	return &Repository{
		queries: queries,
	}
}

func (r *Repository) GetEquation(ctx context.Context, id ID) (*Equation, error) {
	eqn, err := r.queries.GetEquation(ctx, int64(id))
	if err != nil {
		return nil, fmt.Errorf("error reading from database: %w", err)
	}
	return &Equation{
		Id: ID(eqn.ID), 
		Value: eqn.Value,
	}, nil
}

func (r *Repository) GetEquationFromValue(ctx context.Context, value string) (*Equation, error) {
	eqn, err := r.queries.GetEquationFromValue(ctx, value)
	if err != nil {
		return nil, fmt.Errorf("error reading from database: %w", err)
	}
	return &Equation{
		Id: ID(eqn.ID),
		Value: eqn.Value,
	}, nil
}

func (r *Repository) ListEquations(ctx context.Context) ([]*Equation, error) {
	eqns, err := r.queries.ListEquations(ctx); 
	if err != nil {
		return nil, fmt.Errorf("error reading from database: %w", err)
	}
	var equations []*Equation
	for _, eqn := range eqns {
		equations = append(equations, 
			&Equation{
			Id: ID(eqn.ID), 
			Value: eqn.Value,
			})
	}
	return equations, nil
}

func (r *Repository) CreateEquation(ctx context.Context, value string) (*Equation, error) {
	eqn, err := r.queries.CreateEquation(ctx, value)
	if err != nil {
		return nil, fmt.Errorf("error creating equation: %w", err)
	}
	return &Equation{
		Id: ID(eqn.ID),
		Value: eqn.Value,
	}, nil
}

func (r *Repository) UpdateEquation(ctx context.Context, id ID, value string) (*Equation, error) {
	eqn, err := r.queries.UpdateEquation(ctx, db.UpdateEquationParams{int64(id), value})
	if err != nil {
		return nil, fmt.Errorf("error updating equation: %w", err)
	}
	return &Equation{
		Id: ID(eqn.ID),
		Value: eqn.Value,
	}, nil
}

func (r *Repository) DeleteEquation(ctx context.Context, id ID) error {
	err := r.queries.DeleteEquation(ctx, int64(id))
	if err != nil {
		return fmt.Errorf("error deleting equation: %w", err)
	}
	return nil
}

