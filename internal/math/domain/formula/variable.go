package formula

type Variable struct {
	Id ID `json:"id"`
	ClassificationID uint
	Name       string `json:"name"`
	Vcategory  string `json:"vcategory"`
	Subscripts []rune `json:"subscripts"`
	Arguments []*Argument `json:"arguments"`
	Description string     `json:"description"`
}

