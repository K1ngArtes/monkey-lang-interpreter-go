package ast

import "github.com/K1ngArtes/monkey-lang-interpreter-go/token"

type Node interface {
	// TokenLiteral return the value of the token
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

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
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

// TokenLiteral is the name of the variable in the code
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// Identifier is made to implement Expression for simplicity-sake, not because it is an expression
func (i *Identifier) expressionNode() {}
