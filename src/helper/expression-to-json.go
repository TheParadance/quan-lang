package helper

import "theparadance.com/quan-lang/src/expression"

func ExpressionToJson(expr *[]expression.Expr) []map[string]interface{} {
	var jsondata = make([]map[string]interface{}, len(*expr))
	for index, e := range *expr {
		converted := convert(&e)
		jsondata[index] = *converted
	}
	return jsondata
}

func convert(expr *expression.Expr) *map[string]interface{} {
	var jsondata = make(map[string]interface{})
	switch e := (*expr).(type) {
	case expression.AssignExpr:
		jsondata = map[string]interface{}{
			"type":   "AssignExpr",
			"target": e.Target,
			"value":  e.Value,
		}
		jsondata["target"] = convert(&e.Target)
		jsondata["value"] = convert(&e.Value)
	case expression.FuncDef:
		jsondata = map[string]interface{}{
			"type":   "FuncDef",
			"name":   e.Name,
			"params": e.Params,
			"body":   e.Body,
		}
		jsondata["body"] = ExpressionToJson(&e.Body)
	case expression.NumberExpr:
		jsondata = map[string]interface{}{
			"type":  "NumberExpr",
			"value": e.Value,
		}
	case expression.StringExpr:
		jsondata = map[string]interface{}{
			"type":  "StringExpr",
			"value": e.Value,
		}
	case expression.TernaryExpr:
		jsondata = map[string]interface{}{
			"type":       "TernaryExpr",
			"condition":  e.Condition,
			"trueValue":  e.TrueValue,
			"falseValue": e.FalseValue,
		}
		jsondata["trueValue"] = convert(&e.TrueValue)
		jsondata["falseValue"] = convert(&e.FalseValue)
	case expression.VarExpr:
		jsondata = map[string]interface{}{
			"type": "VarExpr",
			"name": e.Name,
		}
	case expression.IfExpr:
		jsondata = map[string]interface{}{
			"type":      "IfExpr",
			"condition": e.Condition,
			"then":      e.Then,
			"else":      e.Else,
		}
		jsondata["then"] = ExpressionToJson(&e.Then)
		jsondata["else"] = ExpressionToJson(&e.Else)
	case expression.BinaryExpr:
		jsondata = map[string]interface{}{
			"type":     "BinaryExpr",
			"left":     e.Left,
			"operator": e.Operator,
			"right":    e.Right,
		}
		jsondata["left"] = convert(&e.Left)
		jsondata["right"] = convert(&e.Right)
	case expression.ReturnExpr:
		jsondata = map[string]interface{}{
			"type":  "ReturnExpr",
			"value": e.Value,
		}
		jsondata["value"] = convert(&e.Value)
	case expression.FuncCall:
		jsondata = map[string]interface{}{
			"type": "FuncCall",
			"name": e.Name,
			"args": e.Args,
		}
		jsondata["args"] = ExpressionToJson(&e.Args)
	case expression.BooleanExpr:
		jsondata = map[string]interface{}{
			"type":  "BooleanExpr",
			"value": e.Value,
		}
	case expression.ObjectExpr:
		jsondata = map[string]interface{}{
			"type":  "ObjectExpr",
			"pairs": e.Pairs,
		}
		d := make(map[string]interface{})
		for key, value := range e.Pairs {
			d[key] = convert(&value)
		}
		jsondata["pairs"] = d
	case expression.MemberExpr:
		jsondata = map[string]interface{}{
			"type":     "MemberExpr",
			"object":   e.Object,
			"property": e.Property,
		}
		jsondata["object"] = convert(&e.Object)
	case expression.TemplateStringExpr:
		jsondata = map[string]interface{}{
			"type":  "TemplateStringExpr",
			"value": e.Value,
		}
		jsondata["value"] = ExpressionToJson(&e.Value)
	case expression.ArrayExpr:
		jsondata = map[string]interface{}{
			"type":     "ArrayExpr",
			"elements": e.Elements,
		}
		jsondata["elements"] = ExpressionToJson(&e.Elements)
	case expression.IndexExpr:
		jsondata = map[string]interface{}{
			"type":  "IndexExpr",
			"array": e.Array,
			"index": e.Index,
		}
		jsondata["array"] = convert(&e.Array)
		jsondata["index"] = convert(&e.Index)

	default:
		jsondata = map[string]interface{}{
			"type": "Unknown",
		}
	}
	return &jsondata
}
