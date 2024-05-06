package math

import (
	"context"
)

type EquationRepository interface {
	GetEquation(ctx context.Context, id ID) (*Equation, error)
	GetEquationFromValue(ctx context.Context, value string) (*Equation, error)
	ListEquations(ctx context.Context) ([]*Equation, error)
	CreateEquation(ctx context.Context, value string) (*Equation, error)
	UpdateEquation(ctx context.Context, id ID, value string) (*Equation, error)
	DeleteEquation(ctx context.Context, id ID) error
}

type EquationService struct {
	repository EquationRepository
}

func NewEquationService(repository EquationRepository) *EquationService {
	return &EquationService{repository: repository}
}

func (s *EquationService) GetEquation(ctx context.Context, id ID) (*Equation, error) {
	return s.repository.GetEquation(ctx, id)
}

func (s *EquationService) GetEquationFromValue(ctx context.Context, value string) (*Equation, error) {
	return s.repository.GetEquationFromValue(ctx, value)
}

func (s *EquationService) ListEquations(ctx context.Context) ([]*Equation, error) {
	return s.repository.ListEquations(ctx)
}

func (s *EquationService) CreateEquation(ctx context.Context, value string) (*Equation, error) {
	return s.repository.CreateEquation(ctx, value)
}

func (s *EquationService) UpdateEquation(ctx context.Context, id ID, value string) (*Equation, error) {
	return s.repository.UpdateEquation(ctx, id, value)
}

func (s *EquationService) DeleteEquation(ctx context.Context, id ID) error {
	return s.repository.DeleteEquation(ctx, id)
}
