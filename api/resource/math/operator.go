package math

import (
	"gorm.io/gorm"
)

type Operator struct {
	gorm.Model `json:"-" gorm:"-"`
	Value string `json:"value"`
	Description string `json:"description"`
}

type OperatorForm struct {
	Value string `json:"value"`
	Description string `json:"description" form:"required"`
}