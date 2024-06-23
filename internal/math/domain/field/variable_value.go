package field

type VariableValue struct {
	Name     string
	Category string
	Mapper   map[string]IVector
}

func NewVariableValue(str ...string) *VariableValue {
	if len(str) == 0 {
		return &VariableValue{
			Mapper: make(map[string]IVector),
		}
	} else if len(str) == 2 {
		return newVariableValueFull(str[0], str[1])
	}
	return nil
}

func newVariableValueFull(name string, category string) *VariableValue {
	return &VariableValue{
		Name:     name,
		Category: category,
		Mapper:   make(map[string]IVector),
	}
}