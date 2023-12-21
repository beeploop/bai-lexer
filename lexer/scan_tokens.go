package lexer

func (S *Scanner) ScanTokens() []Token {
	for S.IsAtEnd() == false {
		S.start = S.current
		S.ScanToken()
	}

	S.tokens = append(S.tokens, *CreateToken(EOF, "", nil, S.line))

	return S.tokens
}
