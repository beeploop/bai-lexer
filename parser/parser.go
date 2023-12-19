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

	for _, token := range P.tokens {
		fmt.Println("token: ", token)
	}

	return nil
}
