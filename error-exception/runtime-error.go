package errorexception

type RuntimeError struct {
	Message         string `json:"message"`
	ConsoleMessages string `json:"console_messages,omitempty"`
}

func (e *RuntimeError) Error() string {
	return e.Message
}
