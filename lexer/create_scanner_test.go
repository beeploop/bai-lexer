package lexer

import (
	"testing"
)

func TestCreateScanner(t *testing.T) {
	source := "hello"

	scanner := CreateScanner(source)

	if scanner.source != source {
		t.Logf("FAILED, Expected: %v, got: %v\n", source, scanner.source)
		t.Fail()
	}

	if scanner.start != 0 {
		t.Logf("FAILED, Expected: %v, got: %v\n", 0, scanner.start)
		t.Fail()
	}

	if scanner.current != 0 {
		t.Logf("FAILED, Expected: %v, got: %v\n", 0, scanner.current)
		t.Fail()
	}

	if scanner.line != 1 {
		t.Logf("FAILED, Expected: %v, got: %v\n", 1, scanner.line)
		t.Fail()
	}
}
