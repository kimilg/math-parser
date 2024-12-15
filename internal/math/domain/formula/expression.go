package formula

type Expression struct {
	EquationId       ID
	ClassificationID uint
	Category         string
	Elements         []Element
	IsCause          bool
	IsEffect         bool
	Description      string
	IsArgument       bool
}

func (e *Expression) SetVariableConstant(constant *Constant) {
	for _, element := range e.Elements {
		switch el := element.(type) {
		case *Variable:
			el.Constant
		}
	}
}

type Argument struct {
	*Expression
	Seq         uint
	SubCategory string
	Name        string
}

func (e *Expression) GetArgumentKinds() []*ArgumentKind {
	argumentKinds := make(map[ArgumentKind]bool)
	for _, element := range e.Elements {
		switch el := element.(type) {
		case *Variable:
			for _, kind := range el.GetArgumentKinds() {
				argumentKinds[*kind] = true
			}
		case *Expression:
			for _, kind := range el.GetArgumentKinds() {
				argumentKinds[*kind] = true
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
