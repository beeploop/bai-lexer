package lexer

import "fmt"

type Token struct {
	Type    TokenType
	Literal any
	Line    int
}

func CreateToken(t_type TokenType, lexeme string, literal any, line int) *Token {
	return &Token{
		Type:    t_type,
		Literal: literal,
		Line:    line,
	}
}

func (T *Token) ToString() string {
	return fmt.Sprintf("%s %v %d", T.Type, T.Literal, T.Line)
}
