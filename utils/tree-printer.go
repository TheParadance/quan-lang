package utils

import "theparadance.com/quan-lang/src/expression"

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
