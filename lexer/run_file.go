package lexer

import (
	"io"
	"os"
)

func runFile(src string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = runLexer(bytes)
	if err != nil {
		return err
	}

    if hadError {
        os.Exit(1)
    }

	return nil
}
