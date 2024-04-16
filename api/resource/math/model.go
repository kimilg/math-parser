package math

import "gorm.io/gorm"

type Form struct {
	Classification string `json:"classification" form:"required"`
	Equation string `json:"equation" form:"required"`	
}

type Classification struct {
	gorm.Model
	Category string `json:"category" form:"required"`
	Object   string `json:"object" form:"required"`
}

type Classifications []*Classification

type VariableForm struct {
	Classification string `json:"classification" form:"required"`
	Variable string `json:"variable" form:"required"`
	Description string `json:"description" form:"required"`
}

type ConstantForm struct {
	Classification string `json:"classification" form:"required"`
	Constant string `json:"constant" form:"required"`
	Value string `json:"value"`
	Description string `json:"description" form:"required"`
}