package math

import "gorm.io/gorm"

type Classification struct {
	gorm.Model
	Category string
	Object   string
	Variable []Variable
	Constant []Constant
}

type Classifications []*Classification

type Variable struct {
	gorm.Model
	Variable string
	Description string
	ClassificationID uint
}

type Constant struct {
	gorm.Model
	Constant string
	Value string
	Description string
	ClassificationID uint
}





type Form struct {
	Classification string `json:"classification" form:"required"`
	Equation string `json:"equation" form:"required"`
}

type ClassificationForm struct {
	Category string `json:"category" form:"required"`
	Object   string `json:"object" form:"required"`
}

type VariableForm struct {
	Variable string `json:"variable" form:"required"`
	Description string `json:"description" form:"required"`
}

type ConstantForm struct {
	Value string `json:"value"`
	Description string `json:"description" form:"required"`
}