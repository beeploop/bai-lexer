package lexer

import (
	"fmt"

	"github.com/BeepLoop/bai-interpreter/parser"
)

func runLexer(src []byte) error {
	scanner := CreateScanner(string(src))
	tokens := scanner.ScanTokens()

	myparser := parser.CreateParser(tokens)
	expr := myparser.Parse()

    fmt.Println("parsed: ", expr)

	return nil
}
