package math

import (
	"gorm.io/gorm"
	"math-parser/api/resource/math/model/classification"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) List() ([]*classification.Classification, error) {
	classifications := make([]*classification.Classification, 0)
	if err := r.db.Find(&classifications).Error; err != nil {
		return nil, err
	}
	return classifications, nil
}