package lexer

import (
	"testing"
)

func TestCreateScanner(t *testing.T) {
	source := "hello"
	scanner := CreateScanner(source)

	if scanner.source != source {
		t.Logf("test 1 - FAILED, Expected: %v, got: %v\n", source, scanner.source)
		t.Fail()
	} else {
		t.Logf("test 1 - PASSED\n")
	}

	if scanner.start != 0 {
		t.Logf("test 2 - FAILED, Expected: %v, got: %v\n", 0, scanner.start)
		t.Fail()
	} else {
		t.Logf("test 2 - PASSED\n")
	}

	if scanner.current != 0 {
		t.Logf("test 3 - FAILED, Expected: %v, got: %v\n", 0, scanner.current)
		t.Fail()
	} else {
		t.Logf("test 3 - PASSED\n")
	}

	if scanner.line != 1 {
		t.Logf("test 4 - FAILED, Expected: %v, got: %v\n", 1, scanner.line)
		t.Fail()
	} else {
		t.Logf("test 4 - PASSED\n")
	}
}
