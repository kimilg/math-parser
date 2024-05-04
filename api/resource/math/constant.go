package math

import "gorm.io/gorm"

type Constant struct {
	gorm.Model
	Constant string
	Value string
	Description string
	ClassificationID uint
}

type ConstantForm struct {
	Value string `json:"value"`
	Description string `json:"description" form:"required"`
}