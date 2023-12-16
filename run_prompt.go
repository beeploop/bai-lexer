package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func runPrompt() {
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
	}

	fmt.Println("Thank you bye!")
}
