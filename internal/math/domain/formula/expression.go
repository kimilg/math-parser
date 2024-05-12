package formula

type Expression struct {
	EquationId ID
	ClassificationID uint
	Category string
	Elements []Element `json:"elements"`
	IsCause  bool      `json:"isCause"`
	IsEffect bool      `json:"isEffect"`
	Description string `json:"description"`
}

type ExpressionsForm struct {
	Value string `json:"value"`
	Description string `json:"description" form:"required"`
}