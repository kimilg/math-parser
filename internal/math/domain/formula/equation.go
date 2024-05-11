package formula

type ID int64

type Equation struct {
	Id    ID     `json:"id" form:"required"`
	Value string `json:"equation" form:"required"`
}



