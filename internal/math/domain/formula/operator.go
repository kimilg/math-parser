package formula

const (
	PLUS byte = iota
	MINUS
	MULT
	DIV
	POW
	EQUAL
)

type Operator struct {
	Value byte `json:"value"`
	Description string `json:"description"`
}

type OperatorForm struct {
	Value string `json:"value"`
	Description string `json:"description" form:"required"`
}

