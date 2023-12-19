package parser

import (
	"fmt"

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

func (P *Parser) Parse() Expression {

	for P.isAtEnd() == false {
		fmt.Println(P.tokens[P.current])
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
