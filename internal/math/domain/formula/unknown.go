package formula

import (
	"fmt"
	"math-parser/internal/math/domain/field"
)

type Unknown struct {
	Name  string
	Value field.IVector
}

func (u *Unknown) hasValue() bool {
	return u.Value != nil
}

func (u *Unknown) flipSign() *Unknown {
	u.Value.FlipSign()
	return u
}

func (u *Unknown) isSameType(other *Unknown) bool {
	return u.Value.GetDimensions() == other.Value.GetDimensions()
}

func (u *Unknown) Plus(other any) (*Unknown, error) {
	if otherU, ok := other.(*Unknown); ok {
		sum, err := u.Value.Plus(otherU.Value)
		if err != nil {
			return nil, err
		}

		u.Value = sum

		return &Unknown{
			Name:  u.Name + ":" + otherU.Name,
			Value: sum,
		}, nil
	}
	if constant, ok := other.(float64); ok {
		sum, err := u.Value.Plus(field.NewScalar(constant))
		if err != nil {
			return nil, err
		}
		return &Unknown{
			Name:  u.Name,
			Value: sum,
		}, nil
	}

	return nil, fmt.Errorf("plus fail - unknown type")
}

func (u *Unknown) Mult(other any) (*Unknown, error) {
	if o, ok := other.(*Unknown); ok {
		return &Unknown{Name: u.Name + ":" + o.Name,
			Value: u.Value.Mult(o.Value.GetSize()),
		}, nil
	}
	if constant, ok := other.(float64); ok {
		return &Unknown{Name: u.Name,
			Value: u.Value.Mult(constant),
		}, nil
	}
	return nil, fmt.Errorf("cannot multiply to the Unknown: %v", u.Name)
}
