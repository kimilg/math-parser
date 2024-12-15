package formula

type Classification struct {
	Category    string
	Type        string
	IsCause     bool
	IsEffect    bool
	Expressions []*Expression
	Variables   []*Variable
}

type Classifications []*Classification

type ClassificationForm struct {
	Category string
	Object   string
}
