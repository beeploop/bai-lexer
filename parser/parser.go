package parser

import (
	"github.com/BeepLoop/bai-interpreter/types"
)

type Parser struct {
	current int
	tokens  []types.Token
}

func CreateParser(tokens []types.Token) *Parser {
	return &Parser{
		current: 0,
		tokens:  tokens,
	}
}

func (P *Parser) Parse() Expr {

	for P.isAtEnd() == false {
        token := P.tokens[P.current]
        Expression(token)
		P.current++
	}

	return nil
}

func (P *Parser) isAtEnd() bool {
	if P.tokens[P.current].T_type == types.EOF {
		return true
	}

	return false
}
