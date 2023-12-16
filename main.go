package main

import (
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		runPrompt()
	} else {
		runFile(args[1])
	}
}
