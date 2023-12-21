package lexer

import "testing"

func TestPeek(t *testing.T) {

	tests := []struct {
		source   string
		expected byte
	}{
		{source: "", expected: 0},
		{source: "a", expected: 97},
	}

	for i, test := range tests {
		scanner := CreateScanner(test.source)
		peeked := scanner.Peek()

		iter := i + 1
		if peeked != test.expected {
			t.Logf("test %d - FAILED, Expected %v, Got: %v\n", iter, test.expected, peeked)
			t.Fail()
		} else {
			t.Logf("test %d - PASSED\n", iter)
		}
	}

}
