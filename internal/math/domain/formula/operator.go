package formula

const (
	PLUS byte = iota
	MINUS
	MULT
	DIV
	POW
	EQUAL
)

type Operator struct {
	Value       byte
	Description string
}
