package formula

import "math-parser/internal/math/domain/field"

type Result struct {
	Unknowns []Unknown
	Constant float32
}

type Unknown struct {
	Name string
	Coefficient float32
	Value field.IVector
}

func (r *Result) Plus(other *Result) *Result {
	if r.Unknowns == nil && other.Unknowns == nil {
		return &Result{Constant: r.Constant + other.Constant}
	}
	
	unknowns := make([]Unknown, len(r.Unknowns) + len(other.Unknowns))
	unknowns = append(unknowns, r.Unknowns...)
	unknowns = append(unknowns, other.Unknowns...)
	return &Result{Unknowns: unknowns, Constant: r.Constant + other.Constant}
}

func (r *Result) HasSingleUnknown() bool {
	return r.getUnknownNums() == 1
}

func (r *Result) IsConstant() bool {
	return r.getUnknownNums() == 0
}

func (r *Result) GetUnknown() Unknown {
	return r.Unknowns[0]
}

func (r *Result) getUnknownNums() int {
	cnt := 0;
	for _, u := range r.Unknowns {
		if u.Value == nil {
			cnt += 1
		}
	}
	
	return cnt
}