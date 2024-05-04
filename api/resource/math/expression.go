package math

import (
	"gorm.io/gorm"
)

type Expression struct {
	gorm.Model `json:"-" gorm:"-"`
	Element `gorm:"embedded"`
	Elements []Element `json:"elements"`
	Description string `json:"description"`
}

type ExpressionsForm struct {
	Value string `json:"value"`
	Description string `json:"description" form:"required"`
}