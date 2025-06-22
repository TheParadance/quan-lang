package helper

import (
	"theparadance.com/quan-lang/src/token"
)

func TokenToJson(token *[]token.Token) []map[string]interface{} {
	var jsondata = make([]map[string]interface{}, len(*token))
	for index, t := range *token {
		jsondata[index] = map[string]interface{}{
			"type":    t.Type,
			"literal": t.Literal,
			"parts":   TokenToJson(&t.Parts),
		}
	}
	return jsondata
}
