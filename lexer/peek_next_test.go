package lexer

import "testing"

func TestPeekNext(t *testing.T) {
	tests := []struct {
		source   string
		expected byte
	}{
		{source: "", expected: 0},
		{source: "ab", expected: 98},
		{source: "ba", expected: 97},
	}

	for i, test := range tests {
		scanner := CreateScanner(test.source)
		peeked := scanner.PeekNext()

		iter := i + 1
		if peeked != test.expected {
			t.Logf("test %d - FAILED, Exected: %v, Got: %v\n", iter, test.expected, peeked)
			t.Fail()
		} else {
			t.Logf("test %d - PASSED\n", iter)
		}
	}
}
