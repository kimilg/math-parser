package field

const (
	Max Pos = 5000
	PositionMax int64 = 125000000000
)
var MaxPosition Position = Position{Max, Max, Max}

type DFPosition struct {
	DPosition Position
	FPosition Position
}
type Field struct {
	DPosition Position
	Displacement Vector
	FPosition Position
	Force Vector
}
type Fields struct {
	FieldMap map[DFPosition]Field
}

func NewFields() Fields {
	return Fields{
		FieldMap:	make(map[DFPosition]Field),
	}
}