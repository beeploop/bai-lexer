package lexer

type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func CreateScanner(source string) *Scanner {
	return &Scanner{
		source:  source,
		start:   0,
		current: 0,
		line:    1,
	}
}

func (S *Scanner) ScanTokens() []Token {
	for S.isAtEnd() == false {
		S.start = S.current
		S.scanToken()
	}

	S.tokens = append(S.tokens, *CreateToken(EOF, "", nil, S.line))
	return S.tokens
}

func (S *Scanner) scanToken() {
	c := S.Advance()

	switch c {
	case '(':
		S.AddToken(LEFT_PAREN, nil)
	case ')':
		S.AddToken(RIGHT_PAREN, nil)
	case '{':
		S.AddToken(LEFT_BRACE, nil)
	case '}':
		S.AddToken(RIGHT_BRACE, nil)
	case ',':
		S.AddToken(COMMA, nil)
	case '.':
		S.AddToken(DOT, nil)
	case '-':
		S.AddToken(MINUS, nil)
	case '+':
		S.AddToken(PLUS, nil)
	case ';':
		S.AddToken(SEMICOLON, nil)
	case '*':
		S.AddToken(STAR, nil)
	default:
        Error(S.line, "Unexpected character.")
	}
}

func (S *Scanner) Advance() byte {
	S.current++
	return S.source[S.current-1]
}

func (S *Scanner) AddToken(t_type TokenType, literal any) {
	text := S.source[S.start:S.current]
	S.tokens = append(S.tokens, *CreateToken(t_type, text, literal, S.line))
}

func (S *Scanner) isAtEnd() bool {
	res := S.current >= len(S.source)
	return res
}
