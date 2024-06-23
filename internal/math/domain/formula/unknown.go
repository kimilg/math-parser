package formula

import (
	"fmt"
	"math-parser/internal/math/domain/field"
)

type Unknown struct {
	Name string
	Coefficient float32
	Value field.IVector
}

func (u *Unknown) Plus(other any) (*Unknown, error) {
	if otherU, ok := other.(*Unknown); ok {
		val, err := u.Value.Plus(otherU.Value)
		if err != nil {
			return nil, err
		}
		return &Unknown{
			Name:        u.Name + ":" + otherU.Name,
			Coefficient: 1,
			Value:       val.Mult(u.Coefficient).Mult(otherU.Coefficient),
		}, nil
	}
	if constant, ok := other.(float32); ok {
		val, err := u.Value.Plus(field.NewScalar(constant))
		if err != nil {
			return nil, err
		}
		return &Unknown{
			Name: u.Name,
			Coefficient: 1,
			Value: val.Mult(u.Coefficient),
		}, nil
	}
	
	return nil, fmt.Errorf("plus fail - unknown type")
}

func (u *Unknown) Mult(other any) (*Unknown, error) {
	if unknown, ok := other.(*Unknown); ok {
		return u.multUnknown(unknown), nil
	}
	if constant, ok := other.(float32); ok {
		return u.multConstant(constant), nil
	}
	return nil, fmt.Errorf("cannot multiply to the Unknown: %v", u.Name)
}

func (u *Unknown) multUnknown(other *Unknown) *Unknown {
	if u.Value == nil || other.Value == nil {
		fmt.Errorf("unknown value is nil between (%v, %v)", u.Name, other.Name)
	}
	return &Unknown{Name: u.Name + ":" + other.Name,
		Coefficient: u.Coefficient * other.Coefficient,
		Value: nil,
	}
}

func (u *Unknown) multConstant(constant float32) *Unknown {
	if u.Value == nil {
		fmt.Errorf("unknown value is nil (%v)", u.Name)
	}
	return &Unknown{Name: u.Name,
		Coefficient: u.Coefficient * constant,
		Value: u.Value.Mult(constant),
	}
}

