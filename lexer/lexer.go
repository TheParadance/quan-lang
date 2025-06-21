package lexer

import (
	"fmt"
	"unicode"

	"theparadance.com/quan-lang/token"
)

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

func IsDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

func Lex(input string) []token.Token {
	var tokens []token.Token
	i := 0
	for i < len(input) {
		ch := rune(input[i])

		// Skip whitespace
		if unicode.IsSpace(ch) {
			i++
			continue
		}

		// Identifiers or keywords
		if isLetter(ch) {
			start := i
			for i < len(input) && (isLetter(rune(input[i])) || IsDigit(rune(input[i]))) {
				i++
			}
			lit := input[start:i]
			typ := token.TokenIdent
			switch lit {
			case "if":
				typ = token.TokenIf
			case "else":
				typ = token.TokenElse
			case "fn":
				typ = token.TokenFn
			case "return":
				typ = token.TokenReturn
			case "true":
				typ = token.TokenTrue
			case "false":
				typ = token.TokenFalse
			}
			tokens = append(tokens, token.Token{Type: typ, Literal: lit})
			continue
		}

		// Numbers (integer only)
		if IsDigit(ch) {
			start := i
			for i < len(input) && IsDigit(rune(input[i])) {
				i++
			}
			tokens = append(tokens, token.Token{Type: token.TokenNumber, Literal: input[start:i]})
			continue
		}

		// Operators & punctuation
		switch ch {
		case '+':
			tokens = append(tokens, token.Token{Type: token.TokenPlus, Literal: "+"})
			i++
		case '-':
			tokens = append(tokens, token.Token{Type: token.TokenMinus, Literal: "-"})
			i++
		case '*':
			tokens = append(tokens, token.Token{Type: token.TokenStar, Literal: "*"})
			i++
		case '/':
			if i+1 < len(input) && input[i+1] == '/' {
				// Skip comment
				i += 2
				for i < len(input) && input[i] != '\n' {
					i++
				}
			} else {
				tokens = append(tokens, token.Token{Type: token.TokenSlash, Literal: "/"})
				i++
			}
		case '%':
			tokens = append(tokens, token.Token{Type: token.TokenMod, Literal: "%"})
			i++
		case '^':
			tokens = append(tokens, token.Token{Type: token.TokenCaret, Literal: "^"})
			i++
		case '=':
			if i+1 < len(input) && input[i+1] == '=' {
				tokens = append(tokens, token.Token{Type: token.TokenEqual, Literal: "=="})
				i += 2
			} else {
				tokens = append(tokens, token.Token{Type: token.TokenAssign, Literal: "="})
				i++
			}
		case '!':
			if i+1 < len(input) && input[i+1] == '=' {
				tokens = append(tokens, token.Token{Type: token.TokenNE, Literal: "!="})
				i += 2
			} else {
				panic("Unknown token '!'")
			}
		case '<':
			if i+1 < len(input) && input[i+1] == '=' {
				tokens = append(tokens, token.Token{Type: token.TokenLE, Literal: "<="})
				i += 2
			} else {
				tokens = append(tokens, token.Token{Type: token.TokenLT, Literal: "<"})
				i++
			}
		case '>':
			if i+1 < len(input) && input[i+1] == '=' {
				tokens = append(tokens, token.Token{Type: token.TokenGE, Literal: ">="})
				i += 2
			} else {
				tokens = append(tokens, token.Token{Type: token.TokenGT, Literal: ">"})
				i++
			}
		case '(':
			tokens = append(tokens, token.Token{Type: token.TokenLParen, Literal: "("})
			i++
		case ')':
			tokens = append(tokens, token.Token{Type: token.TokenRParen, Literal: ")"})
			i++
		case '{':
			tokens = append(tokens, token.Token{Type: token.TokenLBrace, Literal: "{"})
			i++
		case '}':
			tokens = append(tokens, token.Token{Type: token.TokenRBrace, Literal: "}"})
			i++
		case ',':
			tokens = append(tokens, token.Token{Type: token.TokenComma, Literal: ","})
			i++
		case ';':
			tokens = append(tokens, token.Token{Type: token.TokenSemicolon, Literal: ";"})
			i++
		case '"':
			i++ // Skip the opening quote
			start := i
			for i < len(input) && input[i] != '"' {
				if input[i] == '\\' && i+1 < len(input) {
					i++ // Skip the escape character
				}
				i++
			}
			if i >= len(input) || input[i] != '"' {
				panic("Unterminated string literal")
			}
			tokens = append(tokens, token.Token{Type: token.TokenString, Literal: input[start:i]})
			i++ // Skip the closing quote
		case '\'':
			if i+2 < len(input) && input[i+1] == '\'' && input[i+2] == '\'' {
				// Start triple quote string
				i += 3
				var parts []token.Token
				var buf []rune

				for i < len(input) {
					if i+2 < len(input) && input[i] == '\'' && input[i+1] == '\'' && input[i+2] == '\'' {
						// end of triple quote string
						break
					}

					ch := rune(input[i])
					if ch == '$' && i+1 < len(input) && input[i+1] == '{' {
						// flush buffer as string token
						if len(buf) > 0 {
							parts = append(parts, token.Token{Type: token.TokenString, Literal: string(buf)})
							buf = nil
						}
						i += 2 // skip ${
						exprStart := i
						depth := 1
						for i < len(input) && depth > 0 {
							if input[i] == '{' {
								depth++
							} else if input[i] == '}' {
								depth--
							}
							i++
						}
						if depth != 0 {
							panic("Unclosed ${ in multiline string")
						}
						exprLiteral := input[exprStart : i-1]
						parts = append(parts, token.Token{Type: token.TokenTemplateString, Literal: exprLiteral})
					} else {
						buf = append(buf, ch)
						i++
					}
				}

				if len(buf) > 0 {
					parts = append(parts, token.Token{Type: token.TokenString, Literal: string(buf)})
				}

				i += 3 // skip closing '''

				tokens = append(tokens, token.Token{
					Type:  token.TokenTemplateString,
					Parts: parts,
				})
				continue
			}
		case ':':
			tokens = append(tokens, token.Token{Type: token.TokenColon, Literal: ":"})
			i++
		case '.':
			tokens = append(tokens, token.Token{Type: token.TokenDot, Literal: "."})
			i++
		default:
			panic(fmt.Sprintf("Unknown character: %c", ch))
		}
	}
	tokens = append(tokens, token.Token{Type: token.TokenEOF, Literal: ""})
	return tokens
}
