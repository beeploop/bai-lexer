package lexer

import (
	"log"
	"os"
	"path/filepath"
)

var hadError = false

func Lexer() {
	args := os.Args

	if len(args) < 2 {
		runPrompt()
	} else {
		ext := filepath.Ext(args[1])
		if ext != ".bai" {
			log.Fatal("Invalid file")
		}

		runFile(args[1])
	}
}
