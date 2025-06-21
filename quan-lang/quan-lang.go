package lang

import (
	environment "theparadance.com/quan-lang/src/env"
	errorexception "theparadance.com/quan-lang/src/error-exception"
	interpreter "theparadance.com/quan-lang/src/intepreter"
	lexer "theparadance.com/quan-lang/src/lexer"
	parser "theparadance.com/quan-lang/src/paraser"
	systemconsole "theparadance.com/quan-lang/src/system-console"
)

var (
	DEBUG_MODE   = "DEBUG"
	RELEASE_MODE = "RELEASE"
)

type Mode string

type ExecuationOption struct {
	Mode    string
	Console systemconsole.SystemConsole
}

func NewExecuationOption(console systemconsole.SystemConsole, mode string) *ExecuationOption {
	return &ExecuationOption{
		Mode:    mode,
		Console: console,
	}
}

type ExecuationResult struct {
	Env             *environment.Env
	ConsoleMessages string
}

func Execuate(program string, env *environment.Env, option *ExecuationOption) (ExecuationResult, error) {
	// p := `
	// 	fn fact(n) {
	// 		if (n <= 1) {
	// 			return 1;
	// 		} else {
	// 			return n * fact(n - 1);
	// 		}
	// 	}

	// 	// comment

	// 	b = x;
	// 	y = fact(b);
	// 	z = 10 + y;
	// `

	defer func() {
		if r := recover(); r != nil {
			option.Console.Println("Error:", r.(string))
			var err errorexception.QuanLangEngineError = &errorexception.RuntimeError{
				Message: r.(string),
			}
			panic(err)
		}
	}()

	tokens := lexer.Lex(program)

	p := parser.Parser{Tokens: tokens}
	ast := p.Parse()

	e := environment.NewEnv(env)

	for _, expr := range ast {
		_, _ = interpreter.Eval(expr, e)
	}

	result := ExecuationResult{
		Env:             e,
		ConsoleMessages: option.Console.String(),
	}
	return result, nil
}
