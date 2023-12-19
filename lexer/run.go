package lexer

import (
	"github.com/BeepLoop/bai-interpreter/parser"
)

func run(src []byte) error {
	scanner := CreateScanner(string(src))
	tokens := scanner.ScanTokens()

	parser.Parser(tokens)

	return nil
}
