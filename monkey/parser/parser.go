package parser

import (
	"go-everyday/monkey/ast"
	"go-everyday/monkey/lexer"
	"go-everyday/monkey/token"
)

type Parser struct {
	l *lexer.Lexer // 字句解析器へのポインタ

	curToken  token.Token // 現在のToken
	peekToken token.Token // 次のToken
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}
