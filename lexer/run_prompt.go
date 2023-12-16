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

		err = run(line)
		if err != nil {
			return err
		}
	}

	fmt.Println("Thank you bye!")
	return nil
}
