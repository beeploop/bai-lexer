package lexer

import "testing"

func TestAdvance(t *testing.T) {
	tests := []struct {
		source   string
		expected byte
	}{
		{source: " ", expected: 32},
		{source: "a", expected: 97},
		{source: "1", expected: 49},
	}

	for i, test := range tests {
		scanner := CreateScanner(test.source)
		result := scanner.Advance()

		iter := i + 1
		if result != test.expected {
			t.Logf("test %d - FAILED: Expected: %v, Got: %v\n", iter, test.expected, result)
			t.Fail()
		} else {
			t.Logf("test %d - PASSED\n", iter)
		}
	}
}
