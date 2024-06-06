package formula

type Variable struct {
	Id ID 				  `json:"id,omitempty"`
	ClassificationID uint `json:"classificationId,omitempty"`
	Name       string     `json:"name,omitempty"`
	Vcategory  string     `json:"vcategory,omitempty"`
	Subscripts []rune 	  `json:"subscripts,omitempty"`
	Arguments []*Argument `json:"arguments,omitempty"`
	Description string    `json:"description,omitempty"`
}

type ArgumentKey struct {
	Category    string
	SubCategory string
}

type VariableStringArguments struct {
	
}