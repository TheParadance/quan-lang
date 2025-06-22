package lang

import (
	debuglevel "theparadance.com/quan-lang/src/debug/debug-level"
	environment "theparadance.com/quan-lang/src/env"
	errorexception "theparadance.com/quan-lang/src/error-exception"
	"theparadance.com/quan-lang/src/expression"
	interpreter "theparadance.com/quan-lang/src/intepreter"
	lexer "theparadance.com/quan-lang/src/lexer"
	parser "theparadance.com/quan-lang/src/paraser"
	systemconsole "theparadance.com/quan-lang/src/system-console"
	"theparadance.com/quan-lang/src/token"
	"theparadance.com/quan-lang/utils"
)

var (
	DEBUG_MODE   = "DEBUG"
	RELEASE_MODE = "RELEASE"
)

type Mode string

type ExecuationOption struct {
	Mode       string
	Console    systemconsole.SystemConsole
	DebugLevel []debuglevel.DebugLevel
}

func NewExecuationOption(console systemconsole.SystemConsole, mode string, debugLevel *[]debuglevel.DebugLevel) *ExecuationOption {
	return &ExecuationOption{
		Mode:       mode,
		Console:    console,
		DebugLevel: *debugLevel,
	}
}

type ExecuationResult struct {
	Env             *environment.Env
	ConsoleMessages string
	Tokens          *[]token.Token
	Expression      *[]expression.Expr
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
			option.Console.Println("[Error]: ", r.(string))
			var err errorexception.QuanLangEngineError = &errorexception.RuntimeError{
				Message: r.(string),
			}
			panic(err)
		}
	}()

	if option.Mode == DEBUG_MODE {
		println("Status: Lexing program")
	}
	tokens := lexer.Lex(program)
	if option.Mode == DEBUG_MODE && utils.ArrayItemContain(option.DebugLevel, debuglevel.LEXER_TOKENS) {
		println("========== Lexed Tokens ==========")
		option.Console.Println("Tokens:")
		for _, token := range tokens {
			println(token.Type, token.Literal)
		}
		println("=============================")
	}

	if option.Mode == DEBUG_MODE {
		println("Status: Parsing program")
	}
	p := parser.Parser{Tokens: tokens}
	ast := p.Parse()
	if option.Mode == DEBUG_MODE && utils.ArrayItemContain(option.DebugLevel, debuglevel.LEXER_TOKENS) {
		println("========== AST Tree ==========")
		for _, expr := range ast {
			utils.PrintExpression(expr, 0)
		}
		println("=============================")
	}

	if option.Mode == DEBUG_MODE {
		println("Status: Environment loaded")
	}
	e := environment.NewEnv(env)

	if option.Mode == DEBUG_MODE {
		println("Status: Executing program")
	}
	for _, expr := range ast {
		_, _ = interpreter.Eval(expr, e)
	}

	result := ExecuationResult{
		Env:             e,
		ConsoleMessages: option.Console.String(),
		Tokens:          &tokens,
		Expression:      &ast,
	}
	return result, nil
}
