package formula

type Constant struct {
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Value       string `json:"value"`
	Description string  `json:"description"`
}

type EquationConstants struct {
	Equation  string      `json:"equation" form:"required"`
	Constants []*Constant `json:"constants" form:"required"`
}
