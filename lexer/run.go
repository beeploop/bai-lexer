package lexer

import (
	"github.com/BeepLoop/bai-interpreter/parser"
)

func run(src []byte) error {
	scanner := CreateScanner(string(src))
	tokens := scanner.ScanTokens()

	myparser := parser.CreateParser(tokens)
	myparser.Parse()

	return nil
}
