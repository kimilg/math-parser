package math

import (
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

//func (r *Repository) List() ([]*Equation, error) {
//	equations := make([]*Equation, 0)
//	if err := r.queries.ListEquations().Error; err != nil {
//		return nil, err
//	}
//	return classifications, nil
//}