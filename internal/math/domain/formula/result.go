package formula

type Result struct {
	Unknowns []*Unknown
	Constant float32
}

func (r *Result) Plus(other *Result) *Result {
	unknowns := make([]*Unknown, len(r.Unknowns)+len(other.Unknowns))
	unknowns = append(unknowns, r.Unknowns...)
	unknowns = append(unknowns, other.Unknowns...)
	compress(unknowns, PLUS)
	return &Result{Unknowns: unknowns, Constant: r.Constant + other.Constant}
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

func (r *Result) Minus(other *Result) *Result {
	for _, unknown := range other.Unknowns {
		unknown.Value.FlipSign()
	}
	unknowns := make([]*Unknown, len(r.Unknowns)+len(other.Unknowns))
	unknowns = append(unknowns, r.Unknowns...)
	unknowns = append(unknowns, other.Unknowns...)
	compress(unknowns, PLUS)
	return &Result{Unknowns: unknowns, Constant: r.Constant - other.Constant}
}

func (r *Result) Mult(other *Result) *Result { // 충돌을 인지할 수 있다.
	unknowns := make([]*Unknown, r.numConstants()*other.numConstants())
	for _, left := range r.Unknowns {
		for _, right := range other.Unknowns {
			res, _ := left.Mult(right)
			unknowns = append(unknowns, res)
		}
	}
	compress(unknowns, PLUS)
	return &Result{Unknowns: unknowns, Constant: 0}
}

//func (r *Result) HasSingleUnknown() bool {
//	return r.getUnknownNums() == 1
//}
//
//func (r *Result) IsConstant() bool {
//	return r.getUnknownNums() == 0
//}
//
//func (r *Result) GetUnknown() *Unknown {
//	return r.Unknowns[0]
//}
//
//func (r *Result) getUnknownNums() int {
//	cnt := 0
//	for _, u := range r.Unknowns {
//		if u.Value == nil {
//			cnt += 1
//		}
//	}
//
//	return cnt
//}

//func remove[T any](s []T, i int) []T {
//	s[i] = s[len(s)-1]
//	return s[:len(s)-1]
//}

func (r *Result) hasConstant() bool {
	return r.Constant != 0
}

func (r *Result) numConstants() int {
	if r.hasConstant() {
		return len(r.Unknowns) + 1
	}
	return len(r.Unknowns)
}
