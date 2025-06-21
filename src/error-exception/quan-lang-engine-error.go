package errorexception

type QuanLangEngineError interface {
	GetMessage() string
	Error() string
}
