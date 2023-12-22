package lexer

import "testing"

func TestIsDigit(t *testing.T) {
	tests := []struct {
		source   byte
		expected bool
	}{
		{source: []byte("a")[0], expected: false},
		{source: []byte("1")[0], expected: true},
		{source: []byte("0")[0], expected: true},
		{source: []byte("_")[0], expected: false},
		{source: []byte("+")[0], expected: false},
	}

	for i, test := range tests {
		scanner := CreateScanner("")
		result := scanner.IsDigit(test.source)

		iter := i + 1
		if result != test.expected {
			t.Logf("test %d - FAILED: Expected: %v, Got: %v\n", iter, test.expected, result)
			t.Fail()
		} else {
			t.Logf("test %d - PASSED\n", iter)
		}
	}
}
