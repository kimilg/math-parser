package field

const (
	Max Pos = 5000
	PositionMax int64 = 125000000000
)
var MaxPosition Position = Position{Max, Max, Max}

type Field struct {
	Force map[Position]Force
	Displacement map[Position]Displacement
}

func NewField() Field {
	return Field{
		Force:       make(map[Position]Force, PositionMax),
		Displacement: make(map[Position]Displacement, PositionMax),
	}
}