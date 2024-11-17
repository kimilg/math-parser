package field

type Variable struct {
	Name     string
	Category string
	Value    IVector
	Mapper   map[string]IVector
}

func (v *Variable) getValue(param string) IVector {
	if val, ok := v.Mapper[param]; ok {
		return val
	}
	return nil
}

func NewVariable(str ...string) *Variable {
	if len(str) == 0 {
		return &Variable{
			Mapper: make(map[string]IVector),
		}
	} else if len(str) == 2 {
		return newVariableFull(str[0], str[1])
	}
	return nil
}

func newVariableFull(name string, category string) *Variable {
	return &Variable{
		Name:     name,
		Category: category,
		Mapper:   make(map[string]IVector),
	}
}
