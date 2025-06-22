package main

import (
	"encoding/json"
	"flag"

	lang "theparadance.com/quan-lang/quan-lang"
	"theparadance.com/quan-lang/server"
	"theparadance.com/quan-lang/src/env"
	systemconsole "theparadance.com/quan-lang/src/system-console"
	"theparadance.com/quan-lang/utils"
)

// server version main
func serverMain() {
	server.Init()
}

// binary app main
func main() {

	programPath := flag.String("i", ".", "The program to execute")
	mode := string(*flag.String("mode", lang.DEBUG_MODE, "Execution mode: DEBUG or RELEASE"))
	envs := flag.String("envs", "{}", "Environment variables in JSON format")
	flag.Parse()

	if mode == lang.DEBUG_MODE {
		println("Quan Lang Engine", mode)
	}

	var envJson map[string]interface{}
	err := json.Unmarshal([]byte(*envs), &envJson)
	if err != nil {
		panic(err)
	}
	program, _ := utils.ReadFile(*programPath)

	if mode == lang.DEBUG_MODE {
		println("Running in DEBUG mode")
		println("========== Program ==========")
		println(program)
		println("=============================")
	}

	console := systemconsole.NewVirtualSystemConsole()
	langOptions := lang.NewExecuationOption(console, mode)
	e := &env.Env{
		Vars: map[string]interface{}{},
		Builtin: map[string]env.BuiltinFunc{
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
		},
	}
	result, _ := lang.Execuate(program, e, langOptions)

	if mode == lang.DEBUG_MODE {
		println("========== System Console ==========")
	}
	println(result.ConsoleMessages)
	if mode == lang.DEBUG_MODE {
		println("=============================")
	}
}

func wasmBuild() {

}
