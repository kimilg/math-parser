package formula

type Expression struct {
	EquationId ID	      `json:"equationId,omitempty"`
	ClassificationID uint `json:"classificationId,omitempty"`
	Category string	      `json:"category,omitempty"`
	Elements []Element    `json:"elements,omitempty"`
	IsCause  bool         `json:"isCause,omitempty"`
	IsEffect bool         `json:"isEffect,omitempty"`
	Description string    `json:"description,omitempty"`
	IsArgument  bool      `json:"isArgument,omitempty"`
}

type Argument struct {
	*Expression			`json:"expression,omitempty"`
	Seq uint		    `json:"seq,omitempty"`
	SubCategory string  `json:"subCategory,omitempty"`
}

type ExpressionsForm struct {
	Value string `json:"value"`
	Description string `json:"description" form:"required"`
}