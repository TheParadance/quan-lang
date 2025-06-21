package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Basic
	TokenIdent  TokenType = "IDENT"
	TokenNumber TokenType = "NUMBER"
	TokenEOF    TokenType = "EOF"

	// Keywords
	TokenIf     TokenType = "IF"
	TokenElse   TokenType = "ELSE"
	TokenFn     TokenType = "FN"
	TokenReturn TokenType = "RETURN"

	// Operators and punctuation
	TokenPlus  TokenType = "PLUS"
	TokenMinus TokenType = "MINUS"
	TokenStar  TokenType = "STAR"
	TokenSlash TokenType = "SLASH"
	TokenMod   TokenType = "MOD"
	TokenCaret TokenType = "CARET"

	TokenAssign TokenType = "ASSIGN"

	TokenEqual TokenType = "EQ"
	TokenNE    TokenType = "NE"
	TokenLT    TokenType = "LT"
	TokenGT    TokenType = "GT"
	TokenLE    TokenType = "LE"
	TokenGE    TokenType = "GE"

	TokenLParen    TokenType = "LPAREN"
	TokenRParen    TokenType = "RPAREN"
	TokenLBrace    TokenType = "LBRACE"
	TokenRBrace    TokenType = "RBRACE"
	TokenComma     TokenType = "COMMA"
	TokenSemicolon TokenType = "SEMICOLON"
)
