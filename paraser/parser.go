package parser

import (
	"fmt"
	"strconv"

	"theparadance.com/quan-lang/expression"
	"theparadance.com/quan-lang/token"
)

var precedences = map[token.TokenType]int{
	token.TokenEqual: 1,
	token.TokenNE:    1,
	token.TokenLT:    1,
	token.TokenLE:    1,
	token.TokenGT:    1,
	token.TokenGE:    1,

	token.TokenPlus:  2,
	token.TokenMinus: 2,

	token.TokenStar:  3,
	token.TokenSlash: 3,
	token.TokenMod:   3,

	token.TokenCaret: 4,
}

type Parser struct {
	Tokens []token.Token
	pos    int
}

func (p *Parser) Parse() []expression.Expr {
	var exprs []expression.Expr
	for p.peek().Type != token.TokenEOF {
		exprs = append(exprs, p.parseStatement())
	}
	return exprs
}

func (p *Parser) peek() token.Token {
	if p.pos >= len(p.Tokens) {
		return token.Token{Type: token.TokenEOF}
	}
	return p.Tokens[p.pos]
}

func (p *Parser) advance() token.Token {
	tok := p.peek()
	p.pos++
	return tok
}

func (p *Parser) match(tt ...token.TokenType) bool {
	if p.pos >= len(p.Tokens) {
		return false
	}
	for _, t := range tt {
		if p.Tokens[p.pos].Type == t {
			p.pos++
			return true
		}
	}
	return false
}

func (p *Parser) consume(t token.TokenType) token.Token {
	tok := p.peek()
	if tok.Type != t {
		panic(fmt.Sprintf("Expected token %s, got %s (%s)", t, tok.Type, tok.Literal))
	}
	p.pos++
	return tok
}

func (p *Parser) parseStatement() expression.Expr {
	if p.match(token.TokenFn) {
		return p.parseFunction()
	}
	if p.match(token.TokenIf) {
		return p.parseIf()
	}
	if p.match(token.TokenReturn) {
		val := p.parseExpr()
		p.match(token.TokenSemicolon) // optional semicolon
		return expression.ReturnExpr{Value: val}
	}
	// Assignment or expression
	expr := p.parseExpr()
	if assign, ok := expr.(expression.AssignExpr); ok {
		p.match(token.TokenSemicolon) // optional semicolon
		return assign
	}
	p.match(token.TokenSemicolon) // optional semicolon
	return expr
}

func (p *Parser) parseFunction() expression.Expr {
	name := p.consume(token.TokenIdent).Literal
	p.consume(token.TokenLParen)
	var params []string
	if !p.match(token.TokenRParen) {
		params = append(params, p.consume(token.TokenIdent).Literal)
		for p.match(token.TokenComma) {
			params = append(params, p.consume(token.TokenIdent).Literal)
		}
		p.consume(token.TokenRParen)
	}
	p.consume(token.TokenLBrace)
	body := p.parseBlock()
	p.consume(token.TokenRBrace)
	return expression.FuncDef{Name: name, Params: params, Body: body}
}

func (p *Parser) parseBlock() []expression.Expr {
	var stmts []expression.Expr
	for p.peek().Type != token.TokenRBrace && p.peek().Type != token.TokenEOF {
		stmts = append(stmts, p.parseStatement())
	}
	return stmts
}

func (p *Parser) parseIf() expression.Expr {
	p.consume(token.TokenLParen)
	cond := p.parseExpr()
	p.consume(token.TokenRParen)
	p.consume(token.TokenLBrace)
	thenBlock := p.parseBlock()
	p.consume(token.TokenRBrace)
	var elseBlock []expression.Expr
	if p.match(token.TokenElse) {
		p.consume(token.TokenLBrace)
		elseBlock = p.parseBlock()
		p.consume(token.TokenRBrace)
	}
	return expression.IfExpr{Condition: cond, Then: thenBlock, Else: elseBlock}
}

func (p *Parser) parseExpr() expression.Expr {
	return p.parsePrecedence(0)
}

func (p *Parser) parsePrecedence(minPrec int) expression.Expr {
	left := p.parsePrimary()

	for {
		tok := p.peek()
		prec, ok := precedences[tok.Type]
		if !ok || prec < minPrec {
			break
		}
		op := p.advance()
		right := p.parsePrecedence(prec + 1)
		left = expression.BinaryExpr{Left: left, Operator: op, Right: right}
	}

	// Check for assignment (lowest precedence)
	if ident, ok := left.(expression.VarExpr); ok && p.match(token.TokenAssign) {
		value := p.parseExpr()
		return expression.AssignExpr{Name: ident.Name, Value: value}
	}

	return left
}

func (p *Parser) parsePrimary() expression.Expr {
	tok := p.peek()
	switch tok.Type {
	case token.TokenTrue, token.TokenFalse:
		p.advance()
		return expression.BooleanExpr{Value: tok.Type == token.TokenTrue}
	case token.TokenNumber:
		p.advance()
		v, _ := strconv.Atoi(tok.Literal)
		return expression.NumberExpr{Value: v}
	case token.TokenIdent:
		p.advance()
		// function call or variable?
		if p.match(token.TokenLParen) {
			var args []expression.Expr
			if p.peek().Type != token.TokenRParen {
				args = append(args, p.parseExpr())
				for p.match(token.TokenComma) {
					args = append(args, p.parseExpr())
				}
			}
			p.consume(token.TokenRParen)
			return expression.FuncCall{Name: tok.Literal, Args: args}
		}
		return expression.VarExpr{Name: tok.Literal}
	case token.TokenLParen:
		p.advance()
		expr := p.parseExpr()
		p.consume(token.TokenRParen)
		return expr
	case token.TokenMinus:
		// unary minus
		p.advance()
		right := p.parsePrimary()
		return expression.BinaryExpr{Left: expression.NumberExpr{Value: 0}, Operator: token.Token{Type: token.TokenMinus}, Right: right}
	default:
		panic("Unexpected token: " + tok.Literal)
	}
}
