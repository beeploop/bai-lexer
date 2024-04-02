package main

import (
	"fmt"

	"github.com/BeepLoop/bai-lexer/lexer"
	"github.com/BeepLoop/bai-lexer/scanner"
)

func main() {
	bytes, err := lexer.ReadBytes("input.bai")
	if err != nil {
		panic(err)
	}

	tokens, err := lexer.Tokenize(bytes, scanner.NewScanner())
	if err != nil {
		panic(err)
	}

	for _, token := range tokens {
		fmt.Println(token)
	}
}
