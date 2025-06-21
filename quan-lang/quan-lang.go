package lang

import (
	environment "theparadance.com/quan-lang/env"
	errorexception "theparadance.com/quan-lang/error-exception"
	interpreter "theparadance.com/quan-lang/intepreter"
	lexer "theparadance.com/quan-lang/lexer"
	parser "theparadance.com/quan-lang/paraser"
)

var (
	DEBUG   = "DEBUG"
	RELEASE = "RELEASE"
)

type Mode string

type ExecuationOption struct {
	Mode string
}

func Execuate(program string, env *environment.Env, option *ExecuationOption) (*environment.Env, error) {
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

	return e, nil
}
