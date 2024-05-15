package formula

import "context"

type Repository interface {
	Get(ctx context.Context, id ID) (*Equation, error)
	GetFromValue(ctx context.Context, value string) (*Equation, error)
	List(ctx context.Context) ([]*Equation, error)
	Insert(ctx context.Context, equation *Equation) (*Equation, error)
	Update(ctx context.Context, equation *Equation) (*Equation, error)
	Delete(ctx context.Context, id ID) error
}