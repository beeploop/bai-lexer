package lexer

import "fmt"

func Error(line int, message string) error {
	report(line, "", message)
	return nil
}

func report(line int, where, message string) error {
	err := fmt.Errorf("[Line %d] Error %s : %s", line, where, message)
	fmt.Println(err)

	return nil
}
