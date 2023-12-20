package parser

import (
	"github.com/BeepLoop/bai-interpreter/types"
)

var curr int
var tokenList []types.Token

type Parser struct {
	current int
	tokens  []types.Token
}

func CreateParser(tokens []types.Token) *Parser {
	curr = 0
	tokenList = tokens

	return &Parser{
		current: 0,
		tokens:  tokens,
	}
}

func (P *Parser) Parse() Expr {

	for P.isAtEnd() == false {
        return Expression()
	}

	return nil
}

func (P *Parser) isAtEnd() bool {
	if P.tokens[P.current].T_type == types.EOF {
		return true
	}

	return false
}
