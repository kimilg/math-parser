package math

type Variable struct {
	ClassificationID uint
	Name string `json:"name"`
	Subscripts []rune `json:"subscripts"`
	Arguments []Expression `json:"arguments"`
	Description string `json:"description"`
}

