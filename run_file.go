package main

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

	err = run(bytes)
	if err != nil {
		return err
	}

	return nil
}
