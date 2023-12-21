package lexer

import (
	"reflect"
	"testing"
)

func TestScanTokens(t *testing.T) {

	tests := []struct {
		source   string
		expected []Token
	}{
		{
			source: "ang lalake kay \"gwapo\"",
			expected: []Token{
				{Type: VAR, Literal: nil, Line: 1},
				{Type: IDETIFIER, Literal: nil, Line: 1},
				{Type: EQUAL, Literal: nil, Line: 1},
				{Type: STRING, Literal: "gwapo", Line: 1},
				{Type: EOF, Literal: nil, Line: 1},
			},
		},
		{
			source: "2",
			expected: []Token{
				{Type: NUMBER, Literal: float64(2), Line: 1},
				{Type: EOF, Literal: nil, Line: 1},
			},
		},
		{
			source: "",
			expected: []Token{
				{Type: EOF, Literal: nil, Line: 1},
			},
		},
		{
			source: "2 + 2 == 4",
			expected: []Token{
				{Type: NUMBER, Literal: float64(2), Line: 1},
				{Type: PLUS, Literal: nil, Line: 1},
				{Type: NUMBER, Literal: float64(2), Line: 1},
				{Type: EQUAL_EQUAL, Literal: nil, Line: 1},
				{Type: NUMBER, Literal: float64(4), Line: 1},
				{Type: EOF, Literal: nil, Line: 1},
			},
		},
	}

	for i, test := range tests {
		scanner := CreateScanner(test.source)
		tokens := scanner.ScanTokens()

		iter := i + 1
		if reflect.DeepEqual(tokens, test.expected) == false {
            t.Logf("test %d - FAILED \nExpected: \n\t%v, \nGot: \n\t%v\n", iter, test.expected, tokens)
			t.Fail()
		} else {
            t.Logf("test %d - PASSED\n", iter)
		}
	}
}
