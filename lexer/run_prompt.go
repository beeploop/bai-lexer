package lexer

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func runPrompt() error {
	fmt.Println("=== welcome to Dui REPL ===")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		if string(line) == "exit" {
			break
		}

		err = runLexer(line)
		if err != nil {
			return err
		}

		hadError = false
	}

	fmt.Println("Thank you bye!")
	return nil
}
