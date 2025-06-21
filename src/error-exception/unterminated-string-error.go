package errorexception

type UnTerminatedStringException struct {
	Message         string `json:"message"`
	ConsoleMessages string `json:"console_messages,omitempty"`
}

func (e *UnTerminatedStringException) Error() string {
	return e.Message
}

func (e *UnTerminatedStringException) GetMessage() string {
	return e.Message
}
