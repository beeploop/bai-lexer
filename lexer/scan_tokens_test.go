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

	for _, test := range tests {
		scanner := CreateScanner(test.source)
		tokens := scanner.ScanTokens()

		if reflect.DeepEqual(tokens, test.expected) == false {
			t.Logf("FAILED, Expected: %v, Got %v\n", test.expected, tokens)
			t.Fail()
		}
	}
}
