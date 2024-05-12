package field

type Pos int

type Coord interface {}

type Position struct {
	X, Y, Z Pos
}

type Vector struct {
	X, Y, Z Pos
	Group string
}

