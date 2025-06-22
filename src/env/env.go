package env

import (
	"theparadance.com/quan-lang/src/expression"
)

type BuiltinFunc func(args []any) (any, error)

type Env struct {
	Vars    map[string]interface{}
	Funcs   map[string]expression.FuncDef
	Builtin map[string]BuiltinFunc
	Parent  *Env
}

func NewEnv(parent *Env) *Env {
	return &Env{
		Vars:   make(map[string]interface{}),
		Funcs:  make(map[string]expression.FuncDef),
		Parent: parent,
	}
}

func (env *Env) GetVar(name string) (interface{}, bool) {
	val, ok := env.Vars[name]
	if !ok && env.Parent != nil {
		return env.Parent.GetVar(name)
	}
	return val, ok
}

func (env *Env) SetVar(name string, val interface{}) {
	env.Vars[name] = val
}

func (env *Env) GetFunc(name string) (expression.FuncDef, bool) {
	fn, ok := env.Funcs[name]
	if !ok && env.Parent != nil {
		return env.Parent.GetFunc(name)
	}
	return fn, ok
}

func (env *Env) GetBuiltin(name string) (BuiltinFunc, bool) {
	fn, ok := env.Builtin[name]
	if !ok && env.Parent != nil {
		return env.Parent.GetBuiltin(name)
	}
	return fn, ok
}
