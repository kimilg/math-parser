package math

import (
	"gorm.io/gorm"
)

type Classification struct {
	gorm.Model
	Category string
	Object   string
	Variable []Variable `gorm:"many2many:variable_classifications;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Classifications []*Classification

type ClassificationForm struct {
	Category string `json:"category" form:"required"`
	Object   string `json:"object" form:"required"`
}