package math

import (
	"gorm.io/gorm"
)

type Expression struct {
	gorm.Model `json:"-" gorm:"-"`
	ClassificationID uint
	Elements []Element `json:"elements"`
	IsCause  bool   `json:"isCause"`
	IsEffect bool   `json:"isEffect"`
	Description string `json:"description"`
}

type ExpressionsForm struct {
	Value string `json:"value"`
	Description string `json:"description" form:"required"`
}