package math

type Classification struct {
	Category string `json:"category"`
	Type     string `json:"type"`
	IsCause  bool   `json:"isCause"`
	IsEffect bool   `json:"isEffect"`
	Expressions []Expression `json:"expressions"`
	Variables []Variable `json:"variables"`
}

type Classifications []*Classification

type ClassificationForm struct {
	Category string `json:"category" form:"required"`
	Object   string `json:"object" form:"required"`
}