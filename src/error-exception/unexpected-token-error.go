package errorexception

type UnExpectedTokenError struct {
	Message         string `json:"message"`
	ConsoleMessages string `json:"console_messages,omitempty"`
}

func (e *UnExpectedTokenError) Error() string {
	return e.Message
}

func (e *UnExpectedTokenError) GetMessage() string {
	return e.Message
}
