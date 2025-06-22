package systemconsole

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type SystemConsole interface {
	Println(args ...any)
	Print(args ...any)
	String() string
	Clear()
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
	defer func() {
		if r := recover(); r != nil {
			panic("Error in VirtualSystemConsole.Print")
		}
	}()

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
		case map[string]interface{}:
			json, _ := MapToPrettyJSON(v)
			virtualConsole.builder.WriteString(json)
		case []interface{}:
			strs := []string{}
			for _, item := range v {
				strs = append(strs, fmt.Sprintf("%v", item))
			}
			virtualConsole.builder.WriteString("[" + strings.Join(strs, ", ") + "]")
		default:
			virtualConsole.builder.WriteString(fmt.Sprintf("%v", v))
			// virtualConsole.builder.WriteString(v.(string)) // Assuming all other types can be converted to string
		}
	}
}

func (virtualConsole *VirtualSystemConsole) String() string {
	return virtualConsole.builder.String()
}

func (virtualConsole *VirtualSystemConsole) Clear() {
	virtualConsole.builder.Reset()
}

func MapToPrettyJSON(m map[string]interface{}) (string, error) {
	bytes, err := json.MarshalIndent(m, "", "  ") // indent with 2 spaces
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
