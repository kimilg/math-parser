package math

import (
	"gorm.io/gorm"
)

type Classification struct {
	gorm.Model `json:"-" gorm:"-"`
	Category string `json:"category"`
	Type     string `json:"type"`
	IsCause  bool   `json:"isCause"`
	IsEffect bool   `json:"isEffect"`
	Expressions []Expression `json:"expressions" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Variables []Variable `json:"variables" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Classifications []*Classification

type ClassificationForm struct {
	Category string `json:"category" form:"required"`
	Object   string `json:"object" form:"required"`
}