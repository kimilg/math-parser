package formula

type Variable struct {
	ClassificationID uint
	Name string            `json:"name"`
	VCategory string
	Subscripts []rune      `json:"subscripts"`
	Arguments []*Expression `json:"arguments"`
	Description string     `json:"description"`
	IsArgument  bool
}

