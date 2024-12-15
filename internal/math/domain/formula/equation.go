package formula

type ID int64

type Equation struct {
	Id        ID                  `json:"id" form:"required"`
	Value     string              `json:"value" form:"required"`
	Category  string              `json:"category" form:"required"`
	Variables []*EquationVariable `json:"variables"`
	Cause     string              `json:"cause"`
	Effect    string              `json:"effect"`
}

type EquationVariable struct {
	Name     string `json:"name,omitempty"`
	Category string `json:"category,omitempty"`
}

func (e *Equation) getVariable(name string, category string) *EquationVariable {
	for _, variable := range e.Variables {
		if name == variable.Name && category == variable.Category {
			return variable
		}
	}
	return nil
}
