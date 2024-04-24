package term

import (
	"gorm.io/gorm"
	"math-parser/api/resource/math/mathobj"
	"math-parser/api/resource/math/variable"
)

type Term struct {
	gorm.Model
	Mathobjs []mathobj.Mathobj
	Value string
	Description string
	ClassificationID uint
}

type SingleTerm struct {
	gorm.Model
	Variable variable.Variable
	Value string
	Description string
	ClassificationID uint
}

type Form struct {
	Value string `json:"value"`
	Description string `json:"description" form:"required"`
}