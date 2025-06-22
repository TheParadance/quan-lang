package builtinfunc

import (
	"theparadance.com/quan-lang/src/env"
	systemconsole "theparadance.com/quan-lang/src/system-console"
)

func BuildInFuncs(console systemconsole.SystemConsole) map[string]env.BuiltinFunc {
	return map[string]env.BuiltinFunc{
		"print": func(args []interface{}) interface{} {
			console.Print(args...)
			return nil
		},
		"println": func(args []interface{}) interface{} {
			console.Println(args...)
			return nil
		},
		"type": func(args []interface{}) interface{} {
			switch args[0].(type) {
			case int:
				return "int"
			case float64:
				return "float64"
			case string:
				return "string"
			case bool:
				return "bool"
			case map[string]interface{}:
				return "object"
			case []interface{}:
				return "array"
			default:
				return "unknown"
			}
		},
	}
}
