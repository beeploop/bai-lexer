package parser

import (
	"fmt"

	"github.com/BeepLoop/bai-interpreter/types"
)

type Parser struct {
	current int
	tokens  []types.Token
}

func CreateParser() *Parser {
	return &Parser{
		current: 0,
	}
}

func (P *Parser) Parse(tokens []types.Token) Expression {
	for _, token := range tokens {
		fmt.Println("token: ", token)
	}

	return nil
}
