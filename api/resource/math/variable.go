package math

import (
	"gorm.io/gorm"
)

type VariableDTO struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Subscripts []string `json:"subscripts"`
	Arguments []string `json:"arguments"`
	Description string `json:"description"`
	Classifications string `json:"classifications"`
}

type Variable struct {
	gorm.Model
	Name string
	Subscripts []string
	Arguments []Variable
	Description string
	Classifications []Classification `gorm:"many2many:variable_classifications;foreignKey:VariableID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID"`
}

