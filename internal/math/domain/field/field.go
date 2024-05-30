package field

const (
	Max         Pos   = 500
	PositionMax int64 = 125000000
)

type IVector interface {
	getSize() uint
}

type Position struct {
	X, Y, Z Pos
}

func (p Position) getSize() uint {
	return 3
}

type Vector struct {
	X, Y, Z     Pos
	Category    string
	SubCategory string
}

func (v Vector) getSize() uint {
	return 3
}

type Scalar struct {
	X           Val
	Category    string
	SubCategory string
}

func (s Scalar) getSize() uint {
	return 1
}

var MaxPosition Position = Position{Max, Max, Max}

type FieldMap map[string]IVector

type Variable struct {
	Name        string
	Category    string
	ValueMapper map[string]IVector
}

func NewVariable(str ...string) *Variable {
	if len(str) == 0 {
		return &Variable{
			ValueMapper: make(map[string]IVector),
		}
	} else if len(str) == 2 {
		return newVariableFull(str[0], str[1])
	}
	return nil
}

func newVariableFull(name string, category string) *Variable {
	return &Variable{
		Name:        name,
		Category:    category,
		ValueMapper: make(map[string]IVector),
	}
}
