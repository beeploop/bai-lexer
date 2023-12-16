package lexer

import "fmt"

func run(src []byte) error {
	scanner := CreateScanner(string(src))
	tokens := scanner.ScanTokens()

	for _, token := range tokens {
        fmt.Println("token: ", token)
	}

	return nil
}
