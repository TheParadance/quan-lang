package lang

import (
	environment "theparadance.com/quan-lang/src/env"
	errorexception "theparadance.com/quan-lang/src/error-exception"
	"theparadance.com/quan-lang/src/expression"
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

	if option.Mode == DEBUG_MODE {
		println("Status: Lexing program")
	}
	tokens := lexer.Lex(program)
	if option.Mode == DEBUG_MODE {
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
	if option.Mode == DEBUG_MODE {
		println("========== AST Tree ==========")
		for _, expr := range ast {
			PrintExpression(expr, 0)
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
	}
	return result, nil
}

func printIndent(index int) {
	for i := 0; i < index; i++ {
		print(" ")
	}
	print("|-")
	for i := 0; i < 4; i++ {
		print("-")
	}
}

func PrintExpression(expr expression.Expr, indent int) {
	printIndent(indent)

	switch e := expr.(type) {
	case expression.AssignExpr:
		println("[AssignExpr]: ", e.Target, "=", e.Value)
		PrintExpression(e.Target, indent+4)
		PrintExpression(e.Value, indent+4)
	case expression.FuncDef:
		print("[FuncDef]:", e.Name, "(")
		for i, param := range e.Params {
			if i > 0 {
				print(", ")
			}
			print(param)
		}
		println(")")
		for _, bodyExpr := range e.Body {
			PrintExpression(bodyExpr, indent+4)
		}
	case expression.NumberExpr:
		println("[NumberExpr]:", e.Value)
	case expression.StringExpr:
		println("[StringExpr]:", e.Value)
	case expression.TernaryExpr:
		println("Ternary Condition:", e.Condition)
		println("True Value:", e.TrueValue, "False Value:", e.FalseValue)
		PrintExpression(e.TrueValue, indent+4)
		PrintExpression(e.FalseValue, indent+4)
	case expression.VarExpr:
		println("[VariableExpr]:", e.Name)
	case expression.IfExpr:
		println("If Condition:", e.Condition)
		println("Then Branches:", len(e.Then), "Else Branches:", len(e.Else))
		for _, thenExpr := range e.Then {
			PrintExpression(thenExpr, indent+4)
		}
		for _, elseExpr := range e.Else {
			PrintExpression(elseExpr, indent+4)
		}
	case expression.BinaryExpr:
		println("[BinaryExpr]:", e.Left, e.Operator.Literal, e.Right)
		PrintExpression(e.Left, indent+4)
		PrintExpression(e.Right, indent+4)
	case expression.ReturnExpr:
		println("[ReturnExpr]:", e.Value)
		PrintExpression(e.Value, indent+4)
	case expression.FuncCall:
		println("[FuncCall]:", e.Name, "Args:", len(e.Args))
		for _, arg := range e.Args {
			PrintExpression(arg, indent+4)
		}
	case expression.BooleanExpr:
		println("[BooleanExpr]:", e.Value)
	case expression.ObjectExpr:
		println("[ObjectExpr]:", len(e.Pairs))
		for key, value := range e.Pairs {
			printIndent(indent + 4)
			println("Key:", key, "Value:", value)
			PrintExpression(value, indent+8)
		}
	case expression.MemberExpr:
		println("[MemberExpr]:", e.Object, "Property:", e.Property)
	case expression.TemplateStringExpr:
		println("Template String Expression with parts:", len(e.Value))
		for _, part := range e.Value {
			switch p := part.(type) {
			case expression.StringExpr:
				printIndent(indent)
				println("String Part:", p.Value)
			default:
				PrintExpression(p, indent+4)
			}
		}
	case expression.ArrayExpr:
		println("[ArrayExpr]:", len(e.Elements))
		for _, element := range e.Elements {
			PrintExpression(element, indent+4)
		}

	case expression.IndexExpr:
		println("[IndexExpr]:", e.Array, "Index:", e.Index)
		PrintExpression(e.Array, indent+4)
		PrintExpression(e.Index, indent+4)

	}
}
