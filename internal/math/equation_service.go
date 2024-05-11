package math

import (
	"context"
	"math-parser/internal/math/domain/formula"
)

type EquationRepository interface {
	Get(ctx context.Context, id formula.ID) (*formula.Equation, error)
	GetFromValue(ctx context.Context, value string) (*formula.Equation, error)
	List(ctx context.Context) ([]*formula.Equation, error)
	Create(ctx context.Context, value string) (*formula.Equation, error)
	Update(ctx context.Context, id formula.ID, value string) (*formula.Equation, error)
	Delete(ctx context.Context, id formula.ID) error
}

type EquationService struct {
	repository EquationRepository
}

func NewEquationService(repository EquationRepository) *EquationService {
	return &EquationService{repository: repository}
}

func (s *EquationService) Get(ctx context.Context, id formula.ID) (*formula.Equation, error) {
	return s.repository.Get(ctx, id)
}

func (s *EquationService) GetFromValue(ctx context.Context, value string) (*formula.Equation, error) {
	return s.repository.GetFromValue(ctx, value)
}

func (s *EquationService) List(ctx context.Context) ([]*formula.Equation, error) {
	return s.repository.List(ctx)
}

func (s *EquationService) Create(ctx context.Context, value string) (*formula.Equation, error) {
	return s.repository.Create(ctx, value)
}

func (s *EquationService) Update(ctx context.Context, id formula.ID, value string) (*formula.Equation, error) {
	return s.repository.Update(ctx, id, value)
}

func (s *EquationService) Delete(ctx context.Context, id formula.ID) error {
	return s.repository.Delete(ctx, id)
}
