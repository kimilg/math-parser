package formula

import "context"

type Repository interface {
	Get(ctx context.Context, id ID) (*Equation, error)
	GetFromValue(ctx context.Context, value string) (*Equation, error)
	List(ctx context.Context) ([]*Equation, error)
	Create(ctx context.Context, value string, category string) (*Equation, error)
	Update(ctx context.Context, id ID, value string) (*Equation, error)
	Delete(ctx context.Context, id ID) error
}