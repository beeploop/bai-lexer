package lexer

import "strconv"

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
	case '!':
		if S.match('=') {
			S.AddToken(BANG_EQUAL, nil)
		} else {
			S.AddToken(BANG, nil)
		}
	case '=':
		if S.match('=') {
			S.AddToken(EQUAL_EQUAL, nil)
		} else {
			S.AddToken(EQUAL, nil)
		}
	case '<':
		if S.match('=') {
			S.AddToken(LESS_EQUAL, nil)
		} else {
			S.AddToken(LESS, nil)
		}
	case '>':
		if S.match('=') {
			S.AddToken(GREATER_EQUAL, nil)
		} else {
			S.AddToken(GREATER, nil)
		}
	case '/':
		if S.match('/') {
			// detect a comment
			for S.peek() != '\n' && S.isAtEnd() == false {
				S.Advance()
			}
		} else {
			S.AddToken(SLASH, nil)
		}
	case '"':
		S.string()
	case ' ':
	case '\r':
	case '\t':
	case '\n':
		S.line++
	default:
		if S.isDigit(c) == true {
			S.number()
		} else {
			Error(S.line, "Unexpected character.")
		}
	}
}

func (S *Scanner) number() {
	for S.isDigit(S.peek()) {
		S.Advance()
	}

	if S.peek() == '.' && S.isDigit(S.peekNext()) {
		S.Advance()

		for S.isDigit(S.peek()) {
			S.Advance()
		}
	}

	value, err := strconv.ParseFloat(S.source[S.start:S.current], 32)
	if err != nil {
		Error(S.line, "Cannot parse to number")
	}

	S.AddToken(NUMBER, value)
}

func (S *Scanner) string() {
	for S.peek() != '"' && S.isAtEnd() == false {
		if S.peek() == '\n' {
			S.line++
		}
		S.Advance()
	}

	if S.isAtEnd() {
		Error(S.line, "Unterminated string")
		return
	}

	S.Advance()

	value := S.source[S.start+1 : S.current-1]
	S.AddToken(STRING, value)
}

func (S *Scanner) match(expected byte) bool {
	if S.isAtEnd() {
		return false
	}
	if S.source[S.current] != expected {
		return false
	}

	S.current++
	return true
}

func (S *Scanner) peek() byte {
	if S.isAtEnd() {
		return 0
	}
	return S.source[S.current]
}

func (S *Scanner) peekNext() byte {
	if S.current+1 > len(S.source) {
		return 0
	}
	return S.source[S.current+1]
}

func (S *Scanner) isDigit(c byte) bool {
	return c >= 48 && c <= 57
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
