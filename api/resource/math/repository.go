package math

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) List() (Classifications, error) {
	classifications := make([]*Classification, 0)
	if err := r.db.Find(&classifications).Error; err != nil {
		return nil, err
	}
	return classifications, nil
}