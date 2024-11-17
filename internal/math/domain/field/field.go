package field

import (
	"errors"
	"math"
)

const (
	Max         Pos   = 500
	PositionMax int64 = 125000000
)

var (
	Dimension = map[string]int{
		"position": 3,
		"time":     -1,
	}
)

type IVector interface {
	//ZeroVector() IVector
	//UnitVector() IVector

	GetDimensions() uint
	GetSize() float64
	FlipSign()
	Plus(IVector) (IVector, error)
	Mult(float64) IVector
}

type Position struct {
	X, Y, Z Pos
}

func (p *Position) GetSize() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y + p.Z*p.Z))
}
func (p *Position) GetDimensions() uint {
	return 3
}
func (p *Position) FlipSign() {
	p.X *= -1
	p.Y *= -1
	p.Z *= -1
}
func (p *Position) Plus(other IVector) (IVector, error) {
	if otherP, ok := other.(*Position); ok {
		p.X += otherP.X
		p.Y += otherP.Y
		p.Z += otherP.Z
		return p, nil
	}
	return nil, errors.New("not a position")
}
func (p *Position) Mult(v float64) IVector {
	p.X *= Pos(v)
	p.Y *= Pos(v)
	p.Z *= Pos(v)
	return p
}

type Vector struct {
	X, Y, Z     Pos
	Category    string
	SubCategory string
}

func (v *Vector) GetSize() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
}
func (v *Vector) GetDimensions() uint {
	return 3
}
func (v *Vector) FlipSign() {
	v.X *= -1
	v.Y *= -1
	v.Z *= -1
}
func (v *Vector) Plus(other IVector) (IVector, error) {
	if otherV, ok := other.(*Position); ok {
		v.X += otherV.X
		v.Y += otherV.Y
		v.Z += otherV.Z
		return v, nil
	}
	return nil, errors.New("not a vector")
}
func (v *Vector) Mult(val float64) IVector {
	v.X *= Pos(val)
	v.Y *= Pos(val)
	v.Z *= Pos(val)
	return v
}

type Scalar struct {
	X           Val
	Category    string
	SubCategory string
}

func (s *Scalar) GetSize() float64 {
	return float64(s.X)
}
func NewScalar(val float64) *Scalar {
	return &Scalar{X: Val(val)}
}
func (s *Scalar) GetDimensions() uint {
	return 1
}
func (s *Scalar) FlipSign() {
	s.X *= -1
}
func (s *Scalar) Plus(other IVector) (IVector, error) {
	if otherS, ok := other.(*Scalar); ok {
		s.X += otherS.X
		return s, nil
	}
	return nil, errors.New("not a vector")
}
func (s *Scalar) Mult(v float64) IVector {
	s.X *= Val(v)
	return s
}

var MaxPosition Position = Position{Max, Max, Max}

type FieldMap map[string]IVector
