package main

import (
	"fmt"
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

	fmt.Println(string(bytes))

	return nil
}
