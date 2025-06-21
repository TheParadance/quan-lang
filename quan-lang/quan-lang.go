package lang

import (
	environment "theparadance.com/quan-lang/env"
	errorexception "theparadance.com/quan-lang/error-exception"
	interpreter "theparadance.com/quan-lang/intepreter"
	lexer "theparadance.com/quan-lang/lexer"
	parser "theparadance.com/quan-lang/paraser"
	systemconsole "theparadance.com/quan-lang/system-console"
)

var (
	DEBUG   = "DEBUG"
	RELEASE = "RELEASE"
)

type Mode string

type ExecuationOption struct {
	Mode    string
	Console systemconsole.SystemConsole
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
			panic(&errorexception.RuntimeError{
				Message: r.(string),
			})
		}
	}()

	tokens := lexer.Lex(program)
	// array.NewArray(&tokens).ForEach(func(item *token.Token, index int) {
	// 	println("Token:", item.Type, "Literal:", item.Literal)
	// })

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
