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
	Seq 		uint	`json:"seq,omitempty"`
	SubCategory string  `json:"subCategory,omitempty"`
	Name        string  `json:"name,omitempty"`
}

func (e *Expression) GetArgumentKinds() []*ArgumentKind {
	argumentKinds := make(map[ArgumentKind]bool)
	for _, element := range e.Elements {
		switch el := element.(type) {
		case *Variable:
			for _, key := range el.GetArgumentKinds() {
				argumentKinds[*key] = true
			}
		case *Expression:
			for _, key := range el.GetArgumentKinds() {
				argumentKinds[*key] = true
			}
		}
	}
	
	return slice(argumentKinds)
}

func slice(mapper map[ArgumentKind]bool) []*ArgumentKind {
	var slice []*ArgumentKind
	
	for key, exist := range mapper {
		if exist {
			slice = append(slice, &key)
		}	
	}
	return slice
}

func (e *Expression) GetArgumentNames() []string {
	names := make(map[string]bool)
	for _, element := range e.Elements {
		switch e := element.(type) {
		case *Variable:
			for _, name := range e.GetArgumentNames() {
				names[name] = true
			}
		case *Expression:
			for _, name := range e.GetArgumentNames() {
				names[name] = true
			}
		}
	}

	var nameSlice []string
	for name, exist := range names {
		if exist {
			nameSlice = append(nameSlice, name)
		}
	}
	return nameSlice
}