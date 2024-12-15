package field

type VariableMapper struct {
	Mapper map[string]*Variable
}

func NewVariableMapper() *VariableMapper {
	var variableMapper VariableMapper
	variableMapper.Mapper = make(map[string]*Variable)
	return &variableMapper
}

func (v *VariableMapper) Put(name string, category string, constant IVector) {
	v.Mapper[name] = NewVariable(name, category, constant)
}

type Variable struct {
	Name     string
	Category string
	Constant IVector
	Value    map[string]IVector
}

func (v *Variable) getValue(param string) IVector {
	if val, ok := v.Value[param]; ok {
		return val
	}
	return nil
}

func NewVariable(param ...any) *Variable {
	if len(param) == 0 {
		return &Variable{
			Value: make(map[string]IVector),
		}
	} else if len(param) == 3 {
		return newVariableFull(param[0].(string), param[1].(string), param[2].(IVector))
	}
	return nil
}

func newVariableFull(name string, category string, constant IVector) *Variable {
	return &Variable{
		Name:     name,
		Category: category,
		Constant: constant,
		Value:    make(map[string]IVector),
	}
}
