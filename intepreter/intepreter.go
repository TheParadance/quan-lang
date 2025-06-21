package interpreter

import (
	"fmt"
	"math"
	"strings"

	environment "theparadance.com/quan-lang/env"
	"theparadance.com/quan-lang/expression"
	"theparadance.com/quan-lang/helper"
	"theparadance.com/quan-lang/token"
)

func Eval(expr expression.Expr, env *environment.Env) (interface{}, bool) {
	switch e := expr.(type) {
	case expression.NumberExpr:
		return e.Value, false
	case expression.StringExpr:
		return e.Value, false
	case expression.TemplateStringExpr:
		// Template string evaluation
		var builder strings.Builder
		for _, part := range e.Value {
			switch expr := part.(type) {
			case expression.StringExpr:
				builder.WriteString(expr.Value)
			default:
				val, _ := Eval(expr, env)
				builder.WriteString(fmt.Sprint(val))
			}
		}
		return builder.String(), false
	case expression.BooleanExpr:
		return e.Value, false
	case expression.VarExpr:
		val, ok := env.GetVar(e.Name)
		if !ok {
			panic("Undefined variable: " + e.Name)
		}
		return val, false
	case expression.AssignExpr:
		val, _ := Eval(e.Value, env)
		switch target := e.Target.(type) {
		case expression.VarExpr:
			env.SetVar(target.Name, val)
		case expression.MemberExpr:
			objVal, _ := Eval(target.Object, env)
			if objMap, ok := objVal.(map[string]interface{}); ok {
				objMap[target.Property] = val
			} else {
				panic("Attempt to assign to property on non-object")
			}
		default:
			panic("Invalid assignment target")
		}
		return val, false
	case expression.BinaryExpr:
		leftVal, _ := Eval(e.Left, env)
		rightVal, _ := Eval(e.Right, env)

		// Helper function: convert interface{} to float64 if possible
		toFloat := func(v interface{}) (float64, bool) {
			switch n := v.(type) {
			case int:
				return float64(n), true
			case float64:
				return n, true
			default:
				return 0, false
			}
		}

		// Helper function: convert interface{} to int if possible
		toInt := func(v interface{}) (int, bool) {
			switch n := v.(type) {
			case int:
				return n, true
			case float64:
				// Only convert if float64 is integral
				if n == float64(int(n)) {
					return int(n), true
				}
				return 0, false
			default:
				return 0, false
			}
		}

		switch e.Operator.Type {
		case token.TokenPlus:
			// String concatenation
			if ls, ok := leftVal.(string); ok {
				if rs, ok := rightVal.(string); ok {
					return ls + rs, false
				}
			}

			// Fallback to numeric addition
			lf, lok := toFloat(leftVal)
			rf, rok := toFloat(rightVal)
			if !lok || !rok {
				panic("Plus operator requires both numeric or both string operands")
			}
			return lf + rf, false
		case token.TokenMinus, token.TokenStar, token.TokenSlash, token.TokenCaret:
			// Arithmetic operators
			lf, lok := toFloat(leftVal)
			rf, rok := toFloat(rightVal)
			if !lok || !rok {
				panic("Arithmetic operators require numeric types")
			}

			switch e.Operator.Type {
			case token.TokenMinus:
				return lf - rf, false
			case token.TokenStar:
				return lf * rf, false
			case token.TokenSlash:
				if rf == 0 {
					panic("Division by zero")
				}
				return lf / rf, false
			case token.TokenCaret:
				return math.Pow(lf, rf), false
			}

		case token.TokenMod:
			// Modulus only for integers
			li, lok := toInt(leftVal)
			ri, rok := toInt(rightVal)
			if !lok || !rok {
				panic("Modulo operator requires integer operands")
			}
			if ri == 0 {
				panic("Modulo by zero")
			}
			return li % ri, false

		case token.TokenEqual, token.TokenNE, token.TokenLT, token.TokenLE, token.TokenGT, token.TokenGE:
			// Equality & Comparison - support int, float64, string, bool

			switch l := leftVal.(type) {
			case int:
				switch r := rightVal.(type) {
				case int:
					return helper.CompareInts(l, r, e.Operator.Type), false
				case float64:
					return helper.CompareFloats(float64(l), r, e.Operator.Type), false
				default:
					panic("Type mismatch in comparison")
				}
			case float64:
				switch r := rightVal.(type) {
				case int:
					return helper.CompareFloats(l, float64(r), e.Operator.Type), false
				case float64:
					return helper.CompareFloats(l, r, e.Operator.Type), false
				default:
					panic("Type mismatch in comparison")
				}
			case string:
				rs, ok := rightVal.(string)
				if !ok {
					panic("Type mismatch in comparison")
				}
				return helper.CompareStrings(l, rs, e.Operator.Type), false
			case bool:
				rb, ok := rightVal.(bool)
				if !ok {
					panic("Type mismatch in comparison")
				}
				return helper.CompareBools(l, rb, e.Operator.Type), false
			default:
				panic("Unsupported type for comparison")
			}

		default:
			panic("Unknown operator: " + e.Operator.Literal)
		}
	case expression.IfExpr:
		cond, _ := Eval(e.Condition, env)

		// Handle bool
		if b, ok := cond.(bool); ok {
			if b {
				for _, stmt := range e.Then {
					val, ret := Eval(stmt, env)
					if ret {
						return val, ret
					}
				}
			} else {
				for _, stmt := range e.Else {
					val, ret := Eval(stmt, env)
					if ret {
						return val, ret
					}
				}
			}
			return nil, false
		}

		// fallback if condition is int
		if i, ok := cond.(int); ok && i != 0 {
			for _, stmt := range e.Then {
				val, ret := Eval(stmt, env)
				if ret {
					return val, ret
				}
			}
		} else {
			for _, stmt := range e.Else {
				val, ret := Eval(stmt, env)
				if ret {
					return val, ret
				}
			}
		}
		return nil, false
	case expression.FuncDef:
		env.Funcs[e.Name] = e
		return 0, false
	case expression.FuncCall:
		// 1. Try user-defined function
		if fn, ok := env.GetFunc(e.Name); ok {
			if len(e.Args) != len(fn.Params) {
				panic(fmt.Sprintf("Function %s expects %d args, got %d", e.Name, len(fn.Params), len(e.Args)))
			}
			// Prepare local environment
			localEnv := environment.NewEnv(env)
			for i, param := range fn.Params {
				argVal, _ := Eval(e.Args[i], env)
				localEnv.SetVar(param, argVal)
			}

			for _, stmt := range fn.Body {
				val, ret := Eval(stmt, localEnv)
				if ret {
					return val, true
				}
			}
			return 0, false
		}

		// 2. Try built-in function
		if builtin, ok := env.GetBuiltin(e.Name); ok {
			var args []interface{}
			for _, argExpr := range e.Args {
				argVal, _ := Eval(argExpr, env)
				args = append(args, argVal)
			}
			return builtin(args), false
		}
	case expression.ReturnExpr:
		val, _ := Eval(e.Value, env)
		return val, true
	case expression.ObjectExpr:
		obj := make(map[string]interface{})
		for k, vExpr := range e.Pairs {
			v, _ := Eval(vExpr, env)
			obj[k] = v
		}
		return obj, false
	// object Member access evaluation: a.x
	case expression.MemberExpr:
		objVal, _ := Eval(e.Object, env)
		if objMap, ok := objVal.(map[string]interface{}); ok {
			return objMap[e.Property], false
		}
		panic("Attempt to access property on non-object")
	default:
		panic("Unknown expression type")
	}

	return 0, false
}
