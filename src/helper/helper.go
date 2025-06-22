package helper

import (
	"theparadance.com/quan-lang/src/object"
	"theparadance.com/quan-lang/src/token"
)

func CompareInts(a, b int, op token.TokenType) int {
	switch op {
	case token.TokenEqual:
		if a == b {
			return 1
		}
	case token.TokenNE:
		if a != b {
			return 1
		}
	case token.TokenLT:
		if a < b {
			return 1
		}
	case token.TokenLE:
		if a <= b {
			return 1
		}
	case token.TokenGT:
		if a > b {
			return 1
		}
	case token.TokenGE:
		if a >= b {
			return 1
		}
	}
	return 0
}

func CompareFloats(a, b float64, op token.TokenType) int {
	switch op {
	case token.TokenEqual:
		if a == b {
			return 1
		}
	case token.TokenNE:
		if a != b {
			return 1
		}
	case token.TokenLT:
		if a < b {
			return 1
		}
	case token.TokenLE:
		if a <= b {
			return 1
		}
	case token.TokenGT:
		if a > b {
			return 1
		}
	case token.TokenGE:
		if a >= b {
			return 1
		}
	}
	return 0
}

func CompareStrings(a, b string, op token.TokenType) int {
	switch op {
	case token.TokenEqual:
		if a == b {
			return 1
		}
	case token.TokenNE:
		if a != b {
			return 1
		}
	default:
		panic("Unsupported string comparison operator")
	}
	return 0
}

func CompareBools(a, b bool, op token.TokenType) int {
	switch op {
	case token.TokenEqual:
		if a == b {
			return 1
		}
	case token.TokenNE:
		if a != b {
			return 1
		}
	default:
		panic("Unsupported bool comparison operator")
	}
	return 0
}

func CompareNulls(a, b *object.Null, op token.TokenType) int {
	switch op {
	case token.TokenEqual:
		if a == b {
			return 1
		}
	case token.TokenNE:
		if a != b {
			return 1
		}
	default:
		panic("Unsupported bool comparison operator")
	}
	return 0
}
