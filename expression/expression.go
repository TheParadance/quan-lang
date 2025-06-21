package expression

import "theparadance.com/quan-lang/token"

type Expr interface{}

type NumberExpr struct {
	Value int
}

type StringExpr struct {
	Value string
}

type TemplateStringExpr struct {
	Value []Expr // mix of StringLiteral and any other expression
}

type VarExpr struct {
	Name string
}

type AssignExpr struct {
	Name  string
	Value Expr
}

type BinaryExpr struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

type IfExpr struct {
	Condition Expr
	Then      []Expr
	Else      []Expr
}

type FuncDef struct {
	Name   string
	Params []string
	Body   []Expr
}

type FuncCall struct {
	Name string
	Args []Expr
}

type ReturnExpr struct {
	Value Expr
}

type BooleanExpr struct {
	Value bool
}
