package packagebuilder

import (
	"encoding/json"
	"syscall/js"

	lang "theparadance.com/quan-lang/quan-lang"
	builtinfunc "theparadance.com/quan-lang/src/builtin-func"
	debuglevel "theparadance.com/quan-lang/src/debug/debug-level"
	"theparadance.com/quan-lang/src/env"
	"theparadance.com/quan-lang/src/helper"
	systemconsole "theparadance.com/quan-lang/src/system-console"
)

func BuildWasm() {
	c := make(chan struct{}, 0)
	js.Global().Set("execute", js.FuncOf(executeForWasm))
	<-c // Keep the Go WASM module running
}

func executeForWasm(this js.Value, args []js.Value) (exeResult any) {
	obj := args[0]
	mode := obj.Get("mode").String()
	vars := jsObjectToMap(obj.Get("vars"))
	program := obj.Get("program").String()
	debugLvVal := obj.Get("debugLv")

	var debugLevels []debuglevel.DebugLevel
	if debugLvVal.InstanceOf(js.Global().Get("Array")) {
		length := debugLvVal.Length()
		for i := 0; i < length; i++ {
			str := debugLvVal.Index(i).String() // convert each item to string here
			debugLevels = append(debugLevels, debuglevel.DebugLevel(str))
		}
	}

	if mode == lang.DEBUG_MODE {
		println("Quan Lang Engine", mode)
	}

	console := systemconsole.NewVirtualSystemConsole()
	langOptions := lang.NewExecuationOption(console, mode, &debugLevels)
	e := &env.Env{
		Vars:    vars,
		Builtin: builtinfunc.BuildInFuncs(console),
	}

	defer func() {
		if r := recover(); r != nil {
			exeResult = js.ValueOf(map[string]interface{}{
				"message": "Fail to run program",
				"payload": map[string]interface{}{
					"program": program,
					"inputs":  vars,
					"outputs": nil,
					"console": console.String(),
					"tokens":  nil,
					"ast":     nil,
				},
			})
		}
	}()
	result, _ := lang.Execuate(program, e, langOptions)

	tokensResult, _ := json.Marshal(helper.TokenToJson(result.Tokens))
	expressionResult, _ := json.Marshal(helper.ExpressionToJson(result.Expression))

	exeResult = js.ValueOf(map[string]interface{}{
		"message": "Program executed successfully",
		"payload": map[string]interface{}{
			"program": program,
			"inputs":  vars,
			"outputs": filterPrimative(result.Env),
			"console": result.ConsoleMessages,
			"tokens":  string(tokensResult),
			"ast":     string(expressionResult),
		},
	})
	return
}

func filterPrimative(env *env.Env) map[string]interface{} {
	result := make(map[string]interface{})
	for name, item := range env.Vars {
		switch item.(type) {
		case int, float64, string, bool:
			result[name] = item
		}
	}
	return result
}

func jsObjectToMap(obj js.Value) map[string]interface{} {
	result := make(map[string]interface{})
	keys := js.Global().Get("Object").Call("keys", obj)

	for i := 0; i < keys.Length(); i++ {
		key := keys.Index(i).String()
		val := obj.Get(key)

		// Convert js.Value to Go value (basic types)
		var goVal interface{}
		switch val.Type() {
		case js.TypeString:
			goVal = val.String()
		case js.TypeNumber:
			goVal = val.Float()
		case js.TypeBoolean:
			goVal = val.Bool()
		case js.TypeObject:
			if val.InstanceOf(js.Global().Get("Array")) {
				// Handle arrays as []interface{}
				arr := make([]interface{}, val.Length())
				for j := 0; j < val.Length(); j++ {
					arr[j] = convertValue(val.Index(j)) // Use a helper function for recursive conversion
				}
				goVal = arr
			} else {
				// Optional: recurse into nested object
				goVal = jsObjectToMap(val)
			}
		default:
			goVal = val.String() // fallback
		}

		result[key] = goVal
	}

	return result
}

func convertValue(val js.Value) interface{} {
	switch val.Type() {
	case js.TypeString:
		return val.String()
	case js.TypeNumber:
		return val.Float()
	case js.TypeBoolean:
		return val.Bool()
	case js.TypeObject:
		if val.InstanceOf(js.Global().Get("Array")) {
			arr := make([]interface{}, val.Length())
			for i := 0; i < val.Length(); i++ {
				arr[i] = convertValue(val.Index(i))
			}
			return arr
		} else {
			return jsObjectToMap(val)
		}
	default:
		return val.String()
	}
}
