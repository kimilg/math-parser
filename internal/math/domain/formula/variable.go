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

type ArgumentKind struct {
	Name		string
	SubCategory string
}

func (v *Variable) GetArgumentKinds() []*ArgumentKind {
	var argumentKinds map[ArgumentKind]bool
	for _, argument := range v.Arguments {
		argumentKinds[ArgumentKind{argument.Name, argument.SubCategory}] = true
	}
	
	return slice(argumentKinds)
}

func (v *Variable) GetArgumentNames() []string {
	var names map[string]bool
	for _, argument := range v.Arguments {
		names[argument.Name] = true
	}
	
	var nameSlice []string
	for name, exist := range names {
		if exist {
			nameSlice = append(nameSlice, name)
		}
	}
	return nameSlice
}

func PopArgumentKind(args []*ArgumentKind) *ArgumentKind {
	if len(args) < 1 {
		return nil
	}
	arg := args[0]
	args = args[1:]
	return arg
}