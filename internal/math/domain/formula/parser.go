package formula

type Parser interface {
	Parse(eq *Equation) interface{}
}
