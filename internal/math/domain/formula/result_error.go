package formula

type ResultError struct {
	Message string
}

func (e *ResultError) Error() string {
	return e.Message
}
