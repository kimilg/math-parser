package terms

import (
	"gorm.io/gorm"
	"math-parser/api/resource/math/model/mathobj"
)

type Terms struct {
	gorm.Model
	Mathobjs []mathobj.Mathobj
	Value string
	Description string
	ClassificationID uint
}

type Form struct {
	Value string `json:"value"`
	Description string `json:"description" form:"required"`
}