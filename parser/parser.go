package parser

import (
	"fmt"

	"github.com/BeepLoop/bai-interpreter/types"
)

func Parser(tokens []types.Token) {
	for _, token := range tokens {
		fmt.Println("token: ", token)
	}

}
