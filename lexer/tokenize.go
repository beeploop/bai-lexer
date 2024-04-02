package lexer

import (
	"github.com/BeepLoop/bai-lexer/scanner"
	"github.com/BeepLoop/bai-lexer/types"
)

func Tokenize(bytes []byte, reader *scanner.Scanner) ([]types.Token, error) {
	tokens := make([]types.Token, 0)

	source := string(bytes)

	for reader.Current >= len(source) == false {
		reader.Start = reader.Current
		char := reader.Advance(source)

		switch char {
		case '(':
			tokens = append(tokens, *types.CreateToken(types.LEFT_PAREN, nil, reader.Line))
		case ')':
			tokens = append(tokens, *types.CreateToken(types.RIGHT_PAREN, nil, reader.Line))
		case '{':
			tokens = append(tokens, *types.CreateToken(types.LEFT_BRACE, nil, reader.Line))
		case '}':
			tokens = append(tokens, *types.CreateToken(types.RIGHT_BRACE, nil, reader.Line))
		case ',':
			tokens = append(tokens, *types.CreateToken(types.COMMA, nil, reader.Line))
		case '.':
			tokens = append(tokens, *types.CreateToken(types.DOT, nil, reader.Line))
		case '-':
			tokens = append(tokens, *types.CreateToken(types.MINUS, nil, reader.Line))
		case '+':
			tokens = append(tokens, *types.CreateToken(types.PLUS, nil, reader.Line))
		case ';':
			tokens = append(tokens, *types.CreateToken(types.SEMICOLON, nil, reader.Line))
		case '*':
			tokens = append(tokens, *types.CreateToken(types.STAR, nil, reader.Line))
		case '!':
			tokens = append(tokens, *types.CreateToken(types.BANG, nil, reader.Line))
		case '=':
			tokens = append(tokens, *types.CreateToken(types.EQUAL, nil, reader.Line))
		case '<':
			tokens = append(tokens, *types.CreateToken(types.LESS, nil, reader.Line))
		case '>':
			tokens = append(tokens, *types.CreateToken(types.GREATER, nil, reader.Line))
		case '/':
			tokens = append(tokens, *types.CreateToken(types.SLASH, nil, reader.Line))
		case '"':
		case ' ':
		case '\r':
		case '\t':
		case '\n':
			reader.Line++
		default:
			Error(reader.Line, "Unsupported character")
		}
	}

	return tokens, nil
}
