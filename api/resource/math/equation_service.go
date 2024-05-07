package math

import (
	"context"
)

type EquationRepository interface {
	Get(ctx context.Context, id ID) (*Equation, error)
	GetFromValue(ctx context.Context, value string) (*Equation, error)
	List(ctx context.Context) ([]*Equation, error)
	Create(ctx context.Context, value string) (*Equation, error)
	Update(ctx context.Context, id ID, value string) (*Equation, error)
	Delete(ctx context.Context, id ID) error
}

type EquationService struct {
	repository EquationRepository
}

func NewEquationService(repository EquationRepository) *EquationService {
	return &EquationService{repository: repository}
}

func (s *EquationService) Get(ctx context.Context, id ID) (*Equation, error) {
	return s.repository.Get(ctx, id)
}

func (s *EquationService) GetFromValue(ctx context.Context, value string) (*Equation, error) {
	return s.repository.GetFromValue(ctx, value)
}

func (s *EquationService) List(ctx context.Context) ([]*Equation, error) {
	return s.repository.List(ctx)
}

func (s *EquationService) Create(ctx context.Context, value string) (*Equation, error) {
	return s.repository.Create(ctx, value)
}

func (s *EquationService) Update(ctx context.Context, id ID, value string) (*Equation, error) {
	return s.repository.Update(ctx, id, value)
}

func (s *EquationService) Delete(ctx context.Context, id ID) error {
	return s.repository.Delete(ctx, id)
}
