package expression

import "theparadance.com/quan-lang/src/token"

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
	Target Expr // could be VarExpr or MemberExpr
	Value  Expr
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

type ObjectExpr struct {
	Pairs map[string]Expr
}

type MemberExpr struct {
	Object   Expr   // e.g. VarExpr{Name: "a"}
	Property string // e.g. "x"
}
