package lexer

import "fmt"

type Token struct {
	T_type    TokenType
	T_lexeme  string
	T_literal any
	T_line    int
}

func CreateToken(t_type TokenType, lexeme string, literal any, line int) *Token {
	return &Token{
		T_type:    t_type,
		T_lexeme:  lexeme,
		T_literal: literal,
		T_line:    line,
	}
}

func (T *Token) ToString() string {
	return fmt.Sprintf("%s %v %d", T.T_type, T.T_literal, T.T_line)
}
