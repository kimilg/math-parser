package classification

import (
	"gorm.io/gorm"
	"math-parser/api/resource/math/model/constant"
	"math-parser/api/resource/math/model/variable"
)

type Classification struct {
	gorm.Model
	Category string
	Object   string
	Variable []variable.Variable
	Constant []constant.Constant
}

type Classifications []*Classification

type Form struct {
	Category string `json:"category" form:"required"`
	Object   string `json:"object" form:"required"`
}