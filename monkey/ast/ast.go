package ast

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statmentNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}
