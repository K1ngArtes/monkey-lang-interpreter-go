package parser

import (
	"github.com/K1ngArtes/monkey-lang-interpreter-go/ast"
	"github.com/K1ngArtes/monkey-lang-interpreter-go/lexer"
	"github.com/K1ngArtes/monkey-lang-interpreter-go/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so curToken and peekToken get set
	l.NextToken()
	l.NextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
