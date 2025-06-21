package systemconsole

import (
	"strconv"
	"strings"
)

type SystemConsole interface {
	Println(args ...any)
	Print(args ...any)
	String() string
}

type VirtualSystemConsole struct {
	builder strings.Builder
}

func NewVirtualSystemConsole() *VirtualSystemConsole {
	return &VirtualSystemConsole{
		builder: strings.Builder{},
	}
}
func (virtualConsole *VirtualSystemConsole) Println(args ...any) {
	virtualConsole.Print(args...)
	virtualConsole.builder.WriteString("\n")
}

func (virtualConsole *VirtualSystemConsole) Print(args ...any) {
	for _, arg := range args {
		switch v := arg.(type) {
		case string:
			virtualConsole.builder.WriteString(v)
		case bool:
			virtualConsole.builder.WriteString(strconv.FormatBool(v))
		case int:
			virtualConsole.builder.WriteString(strconv.Itoa(v))
		case float64:
			virtualConsole.builder.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
		default:
			virtualConsole.builder.WriteString(strconv.Itoa(int(v.(int))))
		}
	}
}

func (virtualConsole *VirtualSystemConsole) String() string {
	return virtualConsole.builder.String()
}
