package math

import (
	"gorm.io/gorm"
)

type Terms struct {
	gorm.Model
	Mathobjs []Mathobj
	Value string
	Description string
	ClassificationID uint
}

type TermsForm struct {
	Value string `json:"value"`
	Description string `json:"description" form:"required"`
}