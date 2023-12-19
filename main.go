package main

import (
	"fmt"
	"os"

	"github.com/BeepLoop/bai-interpreter/lexer"
)

var version = "1.0.0"

func main() {
	args := os.Args

	if len(args) > 1 {

		switch args[1] {
		case "-v":
			fmt.Println(version)
			return

		case "--version":
			fmt.Println(version)
			return
		}
	}

	lexer.Lexer()
}
