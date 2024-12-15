package formula

import (
	"math-parser/internal/math/domain/field"
)

type Result struct {
	Unknowns []*Unknown
	Constant float32
}

func (r *Result) IsEmpty() bool {
	return len(r.Unknowns) == 0
}

func (r *Result) Equal(other *Result) (*Result, error) {
	unknowns := make([]*Unknown, len(r.Unknowns)+len(other.Unknowns))
	unknowns = append(unknowns, r.Unknowns...)
	other.flipSign()
	unknowns = append(unknowns, other.Unknowns...)

	compress(unknowns, PLUS)
	if len(unknowns) != 2 {
		return nil, &ResultError{"more than one unknown found"}
	}

	var name string
	var value field.IVector
	for _, unknown := range unknowns {
		if unknown.hasValue() {
			value = unknown.Value
		} else {
			name = unknown.Name
		}
	}

	if name == "" || value == nil {
		return nil, &ResultError{"unknown value "}
	}

	return &Result{Unknowns: []*Unknown{{name, value}}}, nil
}

func (r *Result) Plus(other *Result) *Result {
	unknowns := make([]*Unknown, len(r.Unknowns)+len(other.Unknowns))
	unknowns = append(unknowns, r.Unknowns...)
	unknowns = append(unknowns, other.Unknowns...)
	compress(unknowns, PLUS)
	return &Result{Unknowns: unknowns}
}

func (r *Result) Minus(other *Result) *Result {
	for _, unknown := range other.Unknowns {
		unknown.Value.FlipSign()
	}
	unknowns := make([]*Unknown, len(r.Unknowns)+len(other.Unknowns))
	unknowns = append(unknowns, r.Unknowns...)
	unknowns = append(unknowns, other.Unknowns...)
	compress(unknowns, PLUS)
	return &Result{Unknowns: unknowns}
}

func (r *Result) Mult(other *Result) *Result { // 충돌을 인지할 수 있다.
	unknowns := make([]*Unknown, len(r.Unknowns)*len(other.Unknowns))
	for _, left := range r.Unknowns {
		for _, right := range other.Unknowns {
			res, _ := left.Mult(right)
			unknowns = append(unknowns, res)
		}
	}
	compress(unknowns, PLUS)
	return &Result{Unknowns: unknowns}
}

func compress(unknowns []*Unknown, operator byte) {
	for i := 0; i < len(unknowns); i++ {
		for j := i + 1; j < len(unknowns); j++ {
			if unknowns[i].isSameType(unknowns[j]) {
				switch operator {
				case PLUS:
					unknowns[i].Plus(unknowns[j])
				}
				unknowns = append(unknowns[:j], unknowns[j+1:]...)
				j--
			}
		}
	}
}

func (r *Result) flipSign() {
	for _, unknown := range r.Unknowns {
		unknown.flipSign()
	}
}

//func remove[T any](s []T, i int) []T {
//	s[i] = s[len(s)-1]
//	return s[:len(s)-1]
//}
