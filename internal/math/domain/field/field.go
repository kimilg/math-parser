package field

const (
	Max         Pos   = 500
	PositionMax int64 = 125000000
)
var (
	LoopSizeMap = map[string]int{
		"position": 3,
		"time": -1,
	}
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
