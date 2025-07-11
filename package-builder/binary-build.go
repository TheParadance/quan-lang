package packagebuilder

import (
	"encoding/json"
	"flag"

	lang "theparadance.com/quan-lang/quan-lang"
	builtinfunc "theparadance.com/quan-lang/src/builtin-func"
	debuglevel "theparadance.com/quan-lang/src/debug/debug-level"
	"theparadance.com/quan-lang/src/env"
	systemconsole "theparadance.com/quan-lang/src/system-console"
	"theparadance.com/quan-lang/utils"
)

// binary app main
func BuildBinary() {

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

	debugLv := []debuglevel.DebugLevel{debuglevel.AST_TREE}
	if mode == lang.DEBUG_MODE && utils.ArrayItemContain(debugLv, debuglevel.PROGRAM) {
		println("Running in DEBUG mode")
		println("========== Program ==========")
		println(program)
		println("=============================")
	}

	console := systemconsole.NewVirtualSystemConsole()
	langOptions := lang.NewExecuationOption(console, mode, &debugLv)
	e := &env.Env{
		Vars: map[string]interface{}{
			"obj": map[string]interface{}{
				"name": "quan",
				"age":  18,
			},
		},
		Builtin: builtinfunc.BuildInFuncs(console),
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
