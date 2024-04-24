package constant

import "gorm.io/gorm"

type Constant struct {
	gorm.Model
	Constant string
	Value string
	Description string
	ClassificationID uint
}

type Form struct {
	Value string `json:"value"`
	Description string `json:"description" form:"required"`
}