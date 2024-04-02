package test

import (
	"reflect"
	"testing"

	"github.com/BeepLoop/bai-lexer/lexer"
	"github.com/BeepLoop/bai-lexer/scanner"
	"github.com/BeepLoop/bai-lexer/types"
)

type tests struct {
	bytes    []byte
	expected []types.Token
}

func TestTokenize(t *testing.T) {
	testCases := []tests{
		{
			bytes: []byte("(){}"),
			expected: []types.Token{
				*types.CreateToken(types.LEFT_PAREN, nil, 1),
				*types.CreateToken(types.RIGHT_PAREN, nil, 1),
				*types.CreateToken(types.LEFT_BRACE, nil, 1),
				*types.CreateToken(types.RIGHT_BRACE, nil, 1),
			},
		},
		{
			bytes: []byte("()\n{}"),
			expected: []types.Token{
				*types.CreateToken(types.LEFT_PAREN, nil, 1),
				*types.CreateToken(types.RIGHT_PAREN, nil, 1),
				*types.CreateToken(types.LEFT_BRACE, nil, 2),
				*types.CreateToken(types.RIGHT_BRACE, nil, 2),
			},
		},
		{
			bytes: []byte("+-=*\n(){}"),
			expected: []types.Token{
				*types.CreateToken(types.PLUS, nil, 1),
				*types.CreateToken(types.MINUS, nil, 1),
				*types.CreateToken(types.EQUAL, nil, 1),
				*types.CreateToken(types.STAR, nil, 1),
				*types.CreateToken(types.LEFT_PAREN, nil, 2),
				*types.CreateToken(types.RIGHT_PAREN, nil, 2),
				*types.CreateToken(types.LEFT_BRACE, nil, 2),
				*types.CreateToken(types.RIGHT_BRACE, nil, 2),
			},
		},
	}

	for _, test := range testCases {
		tokens, err := lexer.Tokenize(test.bytes, scanner.NewScanner())
		if err != nil {
			t.Fatalf(err.Error())
		}

		if reflect.DeepEqual(tokens, test.expected) == false {
			t.Log("Expected: ", test.expected)
			t.Log("Got: ", tokens)
			t.FailNow()
		}
	}
}
