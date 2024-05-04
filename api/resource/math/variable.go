package math

import (
	"gorm.io/gorm"
)

type Variable struct {
	gorm.Model `json:"-" gorm:"-"`
	ClassificationID uint
	Name string `json:"name"`
	Subscripts []rune `json:"subscripts"`
	Arguments []Expression `json:"arguments"`
	Description string `json:"description"`
}

