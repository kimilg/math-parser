package math

import (
	"gorm.io/gorm"
)

type Operator struct {
	gorm.Model
	Value string
	Description string
}

type OperatorForm struct {
	Value string `json:"value"`
	Description string `json:"description" form:"required"`
}