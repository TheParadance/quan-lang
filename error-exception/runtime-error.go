package errorexception

type RuntimeError struct {
	Message string `json:"message"`
}

func (e *RuntimeError) Error() string {
	return e.Message
}
