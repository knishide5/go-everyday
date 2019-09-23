package ast

import "go-everyday/monkey/token"

// ASTのすべてのノードはNodeインターフェースを実装する必要あり
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statmentNode() // dummy
}

type Expression interface {
	Node
	expressionNode() // dummy
}

// ASTのすべてのノードのルート
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	Name *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
