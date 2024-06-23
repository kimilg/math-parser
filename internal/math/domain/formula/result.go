package formula

import "math-parser/internal/math/domain/field"

type Result struct {
	Unknowns []*Unknown
	Constant float32
}


func (r *Result) Plus(other *Result) (*Result, error) {
	var unknown Unknown{"", 1, field.IVector.zeroVector} //static zero vector
	
	for _, left := range r.Unknowns {
		for _, right := range other.Unknowns {
			if left.Value != nil && right.Value != nil {
				res, err := left.Plus(right)
				if err != nil {
					return nil, err
				}
				
			}
		}
	}
	
	
	unknowns := make([]*Unknown, len(r.Unknowns) + len(other.Unknowns))
	unknowns = append(unknowns, r.Unknowns...)
	unknowns = append(unknowns, other.Unknowns...)
	return &Result{Unknowns: unknowns, Constant: r.Constant + other.Constant}
}

func (r *Result) Minus(other *Result) *Result {
	for _, unknown := range other.Unknowns {
		unknown.Value.FlipSign()
	}
	unknowns := make([]*Unknown, len(r.Unknowns) + len(other.Unknowns))
	unknowns = append(unknowns, r.Unknowns...)
	unknowns = append(unknowns, other.Unknowns...)
	return &Result{Unknowns: unknowns, Constant: r.Constant + other.Constant}
}

func (r *Result) Mult(other *Result) *Result { // 충돌을 인지할 수 있다.
	var unknown Unknown	
	for _, left := range r.Unknowns {
		for _, right := range other.Unknowns {
			left.multUnknown(right)
		}
	}
	
	
	unknowns := make([]*Unknown, len(r.Unknowns) + len(other.Unknowns))
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

func (r *Result) GetUnknown() *Unknown {
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