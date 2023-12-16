package lexer

import (
	"os"
)

var hadError = false

func Lexer() {
	args := os.Args

	if len(args) < 2 {
		runPrompt()
	} else {
		runFile(args[1])
	}
}
