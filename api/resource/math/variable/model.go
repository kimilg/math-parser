package variable

import (
	"gorm.io/gorm"
	"math-parser/api/resource/math/mathobj"
)

type Variable struct {
	gorm.Model
	Mathobjs []mathobj.Mathobj
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