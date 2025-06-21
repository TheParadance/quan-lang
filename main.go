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
	mode := flag.String("mode", lang.RELEASE_MODE, "Execution mode: DEBUG or RELEASE")
	envs := flag.String("envs", "{}", "Environment variables in JSON format")
	flag.Parse()

	var envJson map[string]interface{}
	err := json.Unmarshal([]byte(*envs), &envJson)
	if err != nil {
		panic(err)
	}
	program, err := utils.ReadFile(*programPath)

	console := systemconsole.NewVirtualSystemConsole()
	langOptions := lang.NewExecuationOption(console, *mode)
	e := &env.Env{
		Vars: envJson,
		Builtin: map[string]env.BuiltinFunc{
			"print": func(args []interface{}) interface{} {
				console.Print(args...)
				return nil
			},
			"println": func(args []interface{}) interface{} {
				console.Println(args...)
				return nil
			},
		},
	}
	result, err := lang.Execuate(program, e, langOptions)
	println(result.ConsoleMessages)
}

func wasmBuild() {

}
