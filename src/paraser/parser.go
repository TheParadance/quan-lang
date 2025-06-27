package parser

import (
	"fmt"
	"strconv"

	errorexception "theparadance.com/quan-lang/src/error-exception"
	"theparadance.com/quan-lang/src/expression"
	lexer "theparadance.com/quan-lang/src/lexer"
	"theparadance.com/quan-lang/src/token"
)

var precedences = map[token.TokenType]int{
	token.TokenQuestion: 0,

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

func NewParser(tokens []token.Token) *Parser {
	return &Parser{
		Tokens: tokens,
		pos:    0,
	}
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
		panic(errorexception.UnExpectedTokenError{
			Message: fmt.Sprintf("Expected token %s, got %s (%s)", t, tok.Type, tok.Literal),
		})
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
		// Handle `return` with or without a value
		if p.peek().Type == token.TokenSemicolon || p.peek().Type == token.TokenEOF || p.peek().Type == token.TokenRBrace {
			p.match(token.TokenSemicolon) // optional semicolon
			return expression.ReturnExpr{Value: nil}
		}

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

func (p *Parser) parseAnonFunction() expression.Expr {
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
	return expression.FuncDef{
		Params: params,
		Body:   body,
	}
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

		// Handle ternary operator separately, since it's not left-associative
		if tok.Type == token.TokenQuestion && minPrec <= precedences[token.TokenQuestion] {
			p.advance() // consume '?'
			thenExpr := p.parseExpr()
			p.consume(token.TokenColon) // expect ':'
			elseExpr := p.parseExpr()
			left = expression.TernaryExpr{
				Condition:  left,
				TrueValue:  thenExpr,
				FalseValue: elseExpr,
			}
			continue
		}

		prec, ok := precedences[tok.Type]
		if !ok || prec < minPrec {
			break
		}

		op := p.advance()
		right := p.parsePrecedence(prec + 1)
		left = expression.BinaryExpr{
			Left:     left,
			Operator: op,
			Right:    right,
		}
	}

	// Assignment: lowest precedence, parsed after other expressions
	if p.match(token.TokenAssign) {
		value := p.parseExpr()
		switch target := left.(type) {
		case expression.VarExpr:
			return expression.AssignExpr{Target: target, Value: value}
		case expression.MemberExpr:
			return expression.AssignExpr{Target: target, Value: value}
		case expression.IndexExpr: // array a[0] = 4
			return expression.AssignExpr{Target: target, Value: value}
		default:
			panic("Invalid assignment target")
		}
	}

	return left
}

func (p *Parser) parsePrimary() expression.Expr {
	var expr expression.Expr
	tok := p.peek()
	switch tok.Type {
	case token.TokenNull:
		p.advance()
		expr = expression.NullExpr{}
	case token.TokenTemplateString:
		p.advance()
		expr = p.parseTemplateString(tok.Parts)
	case token.TokenTrue, token.TokenFalse:
		p.advance()
		expr = expression.BooleanExpr{Value: tok.Type == token.TokenTrue}
	case token.TokenNumber:
		p.advance()
		v, _ := strconv.Atoi(tok.Literal)
		expr = expression.NumberExpr{Value: float64(v)}
	case token.TokenFloat:
		p.advance()
		v, _ := strconv.ParseFloat(tok.Literal, 64)
		expr = expression.NumberExpr{Value: v}
	case token.TokenString:
		p.advance()
		expr = expression.StringExpr{Value: tok.Literal}
	case token.TokenIdent:
		tok := p.advance()
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
			expr = expression.FuncCall{Name: tok.Literal, Args: args}
		} else {
			expr = expression.VarExpr{Name: tok.Literal}
		}
	case token.TokenLParen:
		p.advance()
		expr := p.parseExpr()
		p.consume(token.TokenRParen)
		return expr
	case token.TokenLBrace:
		expr = p.parseObjectLiteral()
	case token.TokenMinus:
		// unary minus
		p.advance()
		right := p.parsePrimary()
		expr = expression.BinaryExpr{Left: expression.NumberExpr{Value: 0}, Operator: token.Token{Type: token.TokenMinus}, Right: right}
	case token.TokenFn:
		p.advance()
		return p.parseAnonFunction()
	case token.TokenLBracket:
		return p.parseArrayLiteral()
		// default:
		// 	panic("Unexpected token: " + tok.Literal)
	}

	for {
		switch p.peek().Type {
		case token.TokenDot:
			// for object property assignment
			p.advance()
			propTok := p.consume(token.TokenIdent)
			expr = expression.MemberExpr{
				Object:   expr,
				Property: propTok.Literal,
			}

		case token.TokenLBracket:
			// for array index assignment
			p.advance()
			index := p.parseExpr()
			p.consume(token.TokenRBracket)
			expr = expression.IndexExpr{
				Array: expr,
				Index: index,
			}

		case token.TokenLParen:
			// support calling function expressions: (fn(x){...})(5)
			p.advance()
			var args []expression.Expr
			if p.peek().Type != token.TokenRParen {
				args = append(args, p.parseExpr())
				for p.match(token.TokenComma) {
					args = append(args, p.parseExpr())
				}
			}
			p.consume(token.TokenRParen)
			expr = expression.CallExpr{
				Callee: expr,
				Args:   args,
			}
		default:
			return expr
		}
	}
}

func (p *Parser) parseTemplateString(parts []token.Token) expression.Expr {
	var exprParts []expression.Expr

	for _, part := range parts {
		switch part.Type {
		case token.TokenString:
			exprParts = append(exprParts, expression.StringExpr{Value: part.Literal})
		case token.TokenTemplateString: // This represents the embedded ${...}
			sub := NewParserFromString(part.Literal)
			expr := sub.parseExpr()
			exprParts = append(exprParts, expr)
		default:
			panic("Invalid token inside template string: " + part.Type)
		}
	}
	return expression.TemplateStringExpr{Value: exprParts}
}

func (p *Parser) parseObjectLiteral() expression.Expr {
	p.consume(token.TokenLBrace) // consume '{'

	pairs := make(map[string]expression.Expr)

	for p.peek().Type != token.TokenRBrace {
		// Parse key
		var key string
		if p.peek().Type == token.TokenIdent {
			key = p.peek().Literal
			p.advance()
		} else if p.peek().Type == token.TokenString {
			key = p.peek().Literal
			p.advance()
		} else {
			panic("Expected identifier or string as object key")
		}

		p.consume(token.TokenColon) // consume ':'

		// Parse value expression
		value := p.parseExpr()

		pairs[key] = value

		if !p.match(token.TokenComma) {
			break
		}
	}

	p.consume(token.TokenRBrace) // consume '}'

	return expression.ObjectExpr{Pairs: pairs}
}

func (p *Parser) parseArrayLiteral() expression.Expr {
	p.consume(token.TokenLBracket)
	var elements []expression.Expr
	if p.peek().Type != token.TokenRBracket {
		elements = append(elements, p.parseExpr())
		for p.match(token.TokenComma) {
			elements = append(elements, p.parseExpr())
		}
	}
	p.consume(token.TokenRBracket)
	return expression.ArrayExpr{Elements: elements}
}

func NewParserFromString(input string) *Parser {
	tokens := lexer.Lex(input)
	return NewParser(tokens)
}
