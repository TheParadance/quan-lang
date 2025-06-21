package lang

import (
	"theparadance.com/quan-lang/array"
	environment "theparadance.com/quan-lang/env"
	"theparadance.com/quan-lang/expression"
	interpreter "theparadance.com/quan-lang/intepreter"
	lexer "theparadance.com/quan-lang/lexer"
	parser "theparadance.com/quan-lang/paraser"
	"theparadance.com/quan-lang/token"
)

func Execuate(program string, env *environment.Env) (*environment.Env, error) {
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

	println("Executing program:", program)

	tokens := lexer.Lex(program)
	array.NewArray(&tokens).ForEach(func(token *token.Token, index int) {
		println("Token", index, ":", token.Type, "->", token.Literal)
	})

	p := parser.Parser{Tokens: tokens}
	ast := p.Parse()
	array.NewArray(&ast).ForEach(func(expr *expression.Expr, index int) {
		println("AST Node", index, ":", expr)
		switch e := (*expr).(type) {
		case expression.VarExpr:
			println("Variable:", e.Name)
		case expression.AssignExpr:
			println("Assignment:", e.Name, "=", e.Value)
		case expression.NumberExpr:
			println("Number:", e.Value)
		case expression.BinaryExpr:
			println("Binary Expression:", e.Left, e.Operator.Literal, e.Right)
		case expression.FuncDef:
			println("Function:", e.Name, "with params", e.Params)
		case *expression.FuncCall:
			println("Function Call:", e.Name, "with args", e.Args)
		case expression.IfExpr:
			println("If Expression with condition", e.Condition, "and body", e.Then, "else", e.Else)
		case expression.ReturnExpr:
			println("Return Expression with value", e.Value)
		default:
			println("Unknown AST Node Type:", e)
		}
	})

	e := environment.NewEnv(env)

	for _, expr := range ast {
		_, _ = interpreter.Eval(expr, e)
	}

	// fmt.Println("x =", env.Vars["b"])
	// fmt.Println("y =", env.Vars["y"]) // Should print factorial of 5 (120)
	// fmt.Println("z =", env.Vars["z"])

	return e, nil
}
