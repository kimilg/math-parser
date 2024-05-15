package formula

type Expression struct {
	EquationId ID
	ClassificationID uint
	Category string
	SubCategory string
	Elements []Element `json:"elements"`
	IsCause  bool      `json:"isCause"`
	IsEffect bool      `json:"isEffect"`
	Description string `json:"description"`
	IsArgument  bool
	ArgumentSeq uint
}

type ExpressionsForm struct {
	Value string `json:"value"`
	Description string `json:"description" form:"required"`
}