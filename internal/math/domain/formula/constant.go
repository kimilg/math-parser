package formula

type Constant struct {
	Name        string  `json:"name"`
	Value       float32 `json:"value"`
	Description string  `json:"description"`
}

type ConstantForm struct {
	Value       string `json:"value"`
	Description string `json:"description" form:"required"`
}
