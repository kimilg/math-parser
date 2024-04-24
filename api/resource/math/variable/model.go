package variable

import "gorm.io/gorm"

type Variable struct {
	gorm.Model
	Variables []Variable
	value string
	Coefficient int
	Description string
	ClassificationID uint
}

type Form struct {
	Variable string `json:"variable" form:"required"`
	Description string `json:"description" form:"required"`
}