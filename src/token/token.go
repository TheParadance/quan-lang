package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Parts   []Token // For template strings
}

const (
	// Basic
	TokenIdent          TokenType = "IDENT"
	TokenNumber         TokenType = "NUMBER"
	TokenFloat          TokenType = "FLOAT"
	TokenString         TokenType = "STRING"
	TokenTemplateString           = "TEMPLATE_STRING"
	TokenEOF            TokenType = "EOF"

	// Keywords
	TokenIf     TokenType = "IF"
	TokenElse   TokenType = "ELSE"
	TokenFn     TokenType = "FN"
	TokenReturn TokenType = "RETURN"

	// Boolean literals
	TokenTrue  = "TRUE"
	TokenFalse = "FALSE"

	TokenQuestion TokenType = "TERNARY" // For expressions like `condition ? trueValue : falseValue`

	// Operators and punctuation
	TokenPlus  TokenType = "PLUS"
	TokenMinus TokenType = "MINUS"
	TokenStar  TokenType = "STAR"
	TokenSlash TokenType = "SLASH"
	TokenMod   TokenType = "MOD"
	TokenCaret TokenType = "CARET"

	TokenAssign TokenType = "ASSIGN"

	// array
	TokenLBracket TokenType = "LBRACKET" // [
	TokenRBracket TokenType = "RBRACKET" // ]

	// comparator
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
	TokenColon     TokenType = "COLON"
	TokenDot       TokenType = "DOT"
)
