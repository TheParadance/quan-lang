package builtinfunc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"theparadance.com/quan-lang/src/env"
	errorexception "theparadance.com/quan-lang/src/error-exception"
	systemconsole "theparadance.com/quan-lang/src/system-console"
)

func BuildInFuncs(console systemconsole.SystemConsole) map[string]env.BuiltinFunc {
	return map[string]env.BuiltinFunc{
		"print": func(args []interface{}) (interface{}, error) {
			for _, arg := range args {
				switch v := arg.(type) {
				case map[string]interface{}:
					jsonBytes, err := json.Marshal(v)
					if err != nil {
						return nil, err
					}
					console.Println(string(jsonBytes))
				default:
					console.Println(arg)
				}
			}
			return nil, nil
		},
		"println": func(args []interface{}) (interface{}, error) {
			for _, arg := range args {
				switch v := arg.(type) {
				case map[string]interface{}:
					jsonBytes, err := json.Marshal(v)
					if err != nil {
						return nil, err
					}
					console.Println(string(jsonBytes))
				default:
					console.Println(arg)
				}
			}
			return nil, nil
		},
		"type": func(args []interface{}) (interface{}, error) {
			switch args[0].(type) {
			case int:
				return "int", nil
			case float64:
				return "float64", nil
			case string:
				return "string", nil
			case bool:
				return "bool", nil
			case map[string]interface{}:
				return "object", nil
			case []interface{}:
				return "array", nil
			case nil:
				return "null", nil
			default:
				return "unknown", nil
			}
		},
		"string": func(args []interface{}) (interface{}, error) {
			if len(args) != 1 {
				return "error: str() expects 1 argument", &errorexception.RuntimeError{
					Message: "error: str() expects 1 argument",
				}
			}

			switch v := args[0].(type) {
			case int:
				return fmt.Sprintf("%d", v), nil
			case float64:
				return fmt.Sprintf("%g", v), nil
			case bool:
				return fmt.Sprintf("%t", v), nil
			case string:
				return v, nil
			case nil:
				return "null", nil
			case []interface{}:
				elems := make([]string, len(v))
				for i, e := range v {
					elems[i] = fmt.Sprintf("%v", e)
				}
				return "[" + strings.Join(elems, ", ") + "]", nil
			case map[string]interface{}:
				pairs := []string{}
				for key, val := range v {
					pairs = append(pairs, fmt.Sprintf("%s: %v", key, val))
				}
				return "{" + strings.Join(pairs, ", ") + "}", nil
			default:
				return fmt.Sprintf("%v", v), nil
			}
		},
		"int": func(args []interface{}) (interface{}, error) {
			if len(args) != 1 {
				return "error: int() expects 1 argument", &errorexception.RuntimeError{
					Message: "error: int() expects 1 argument",
				}
			}
			switch v := args[0].(type) {
			case int:
				return v, nil
			case float64:
				return int(v), nil
			case bool:
				if v {
					return 1, nil
				} else {
					return 0, nil
				}
			case string:
				if v == "true" {
					return 1, nil
				} else if v == "false" {
					return 0, nil
				}
				f, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return nil, &errorexception.RuntimeError{
						Message: "Value " + fmt.Sprintf("%v", v) + " is not subtype of number",
					}
				}
				return f, nil

			default:
				return nil, &errorexception.RuntimeError{
					Message: "Fail to parse value " + fmt.Sprintf("%v", v) + " to int",
				}
			}
		},
		"float": func(args []interface{}) (interface{}, error) {
			if len(args) != 1 {
				return "error: int() expects 1 argument", &errorexception.RuntimeError{
					Message: "error: int() expects 1 argument",
				}
			}
			switch v := args[0].(type) {
			case int:
				return float64(v), nil
			case float64:
				return v, nil
			case bool:
				if v {
					return 1.0, nil
				} else {
					return 0.0, nil
				}
			case string:
				if v == "true" {
					return 1.0, nil
				} else if v == "false" {
					return 0.0, nil
				}
				f, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return nil, &errorexception.RuntimeError{
						Message: "Value " + fmt.Sprintf("%v", v) + " is not subtype of float",
					}
				}
				return f, nil

			default:
				return nil, &errorexception.RuntimeError{
					Message: "Fail to parse value " + fmt.Sprintf("%v", v) + " to float",
				}
			}
		},
		"bool": func(args []interface{}) (interface{}, error) {
			if len(args) != 1 {
				return "error: int() expects 1 argument", &errorexception.RuntimeError{
					Message: "error: int() expects 1 argument",
				}
			}
			switch v := args[0].(type) {
			case int:
				if v == 0 && v < 1 {
					return false, nil
				} else {
					return true, nil
				}
			case float64:
				if v == 0 && v < 1 {
					return false, nil
				} else {
					return true, nil
				}
			case bool:
				return v, nil
			case string:
				if v == "true" {
					return true, nil
				} else if v == "false" {
					return false, nil
				}
				b, err := strconv.ParseBool(v)
				if err != nil {
					return nil, &errorexception.RuntimeError{
						Message: "Value " + fmt.Sprintf("%v", v) + " is not subtype of bool",
					}
				}
				return b, nil

			default:
				return nil, &errorexception.RuntimeError{
					Message: "Fail to parse value " + fmt.Sprintf("%v", v) + " to float",
				}
			}
		},
		"fetch": func(args []interface{}) (interface{}, error) {
			if len(args) != 1 {
				return nil, &errorexception.RuntimeError{
					Message: "fetch() expects 1 argument (a config object)",
				}
			}

			// Expect argument to be a map (like a JS object)
			config, ok := args[0].(map[string]interface{})
			if !ok {
				return nil, &errorexception.RuntimeError{
					Message: "fetch() argument must be a map (object-like)",
				}
			}

			// Extract URL
			urlVal, ok := config["url"].(string)
			if !ok || urlVal == "" {
				return nil, &errorexception.RuntimeError{
					Message: "fetch() requires a 'url' string field",
				}
			}

			// Method (optional, default GET)
			method := "GET"
			if m, ok := config["method"].(string); ok && m != "" {
				method = m
			}

			// Body (optional)
			var bodyReader io.Reader
			if body, ok := config["body"]; ok {
				switch b := body.(type) {
				case string:
					bodyReader = strings.NewReader(b)
				case []byte:
					bodyReader = bytes.NewReader(b)
				default:
					return nil, &errorexception.RuntimeError{
						Message: "fetch() 'body' must be a string or []byte",
					}
				}
			}

			// Build request
			req, err := http.NewRequest(method, urlVal, bodyReader)
			if err != nil {
				return nil, &errorexception.RuntimeError{
					Message: "fetch() failed to create request: " + err.Error(),
				}
			}

			// Headers (optional)
			if h, ok := config["headers"].(map[string]interface{}); ok {
				for k, v := range h {
					if valStr, ok := v.(string); ok {
						req.Header.Set(k, valStr)
					}
				}
			}

			// Send request
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return nil, &errorexception.RuntimeError{
					Message: "fetch() failed: " + err.Error(),
				}
			}
			defer resp.Body.Close()

			// Read response
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, &errorexception.RuntimeError{
					Message: "fetch() failed reading response: " + err.Error(),
				}
			}

			return string(bodyBytes), nil
		},
		"toMap": func(args []interface{}) (interface{}, error) {
			if len(args) != 1 {
				return nil, &errorexception.RuntimeError{
					Message: "toJson() expects exactly 1 argument: (jsonString)",
				}
			}

			jsonStr, ok := args[0].(string)
			if !ok {
				return nil, &errorexception.RuntimeError{
					Message: "toJson() argument must be a JSON string",
				}
			}

			var result interface{}
			err := json.Unmarshal([]byte(jsonStr), &result)
			if err != nil {
				return nil, &errorexception.RuntimeError{
					Message: "toJson() failed to parse JSON: " + err.Error(),
				}
			}

			return result, nil
		},
	}
}
