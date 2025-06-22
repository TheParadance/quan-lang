package debuglevel

type DebugLevel string

var (
	AST_TREE     DebugLevel = "AST_TREE"
	LEXER_TOKENS DebugLevel = "LEXER_TOKENS"
	PARSER_TREE  DebugLevel = "PARSER_TREE"
	PROGRAM      DebugLevel = "PROGRAM"
)
