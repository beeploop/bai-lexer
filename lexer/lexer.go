package lexer

import (
	"os"
)

func Lexer() {
	args := os.Args

	if len(args) < 2 {
		runPrompt()
	} else {
		runFile(args[1])
	}
}
