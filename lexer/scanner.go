package lexer

import (
	"strconv"

	"github.com/BeepLoop/bai-interpreter/types"
)

var keywords map[string]types.TokenType

type Scanner struct {
	source  string
	tokens  []types.Token
	start   int
	current int
	line    int
}

func CreateScanner(source string) *Scanner {
	keywords = Keywords()

	return &Scanner{
		source:  source,
		start:   0,
		current: 0,
		line:    1,
	}
}

func (S *Scanner) ScanTokens() []types.Token {
	for S.IsAtEnd() == false {
		S.start = S.current
		S.ScanToken()
	}

	S.tokens = append(S.tokens, *types.CreateToken(types.EOF, "", nil, S.line))
	return S.tokens
}

func (S *Scanner) ScanToken() {
	c := S.Advance()

	switch c {
	case '(':
		S.AddToken(types.LEFT_PAREN, nil)
	case ')':
		S.AddToken(types.RIGHT_PAREN, nil)
	case '{':
		S.AddToken(types.LEFT_BRACE, nil)
	case '}':
		S.AddToken(types.RIGHT_BRACE, nil)
	case ',':
		S.AddToken(types.COMMA, nil)
	case '.':
		S.AddToken(types.DOT, nil)
	case '-':
		S.AddToken(types.MINUS, nil)
	case '+':
		S.AddToken(types.PLUS, nil)
	case ';':
		S.AddToken(types.SEMICOLON, nil)
	case '*':
		S.AddToken(types.STAR, nil)
	case '!':
		if S.Match('=') {
			S.AddToken(types.BANG_EQUAL, nil)
		} else {
			S.AddToken(types.BANG, nil)
		}
	case '=':
		if S.Match('=') {
			S.AddToken(types.EQUAL_EQUAL, nil)
		} else {
			S.AddToken(types.EQUAL, nil)
		}
	case '<':
		if S.Match('=') {
			S.AddToken(types.LESS_EQUAL, nil)
		} else {
			S.AddToken(types.LESS, nil)
		}
	case '>':
		if S.Match('=') {
			S.AddToken(types.GREATER_EQUAL, nil)
		} else {
			S.AddToken(types.GREATER, nil)
		}
	case '/':
		if S.Match('/') {
			// detect a comment
			for S.Peek() != '\n' && S.IsAtEnd() == false {
				S.Advance()
			}
		} else {
			S.AddToken(types.SLASH, nil)
		}
	case '"':
		S.String()
	case ' ':
	case '\r':
	case '\t':
	case '\n':
		S.line++
	default:
		if S.IsDigit(c) {
			S.Number()
		} else if S.IsAlpha(c) {
			S.Identifier()
		} else {
			Error(S.line, "Unexpected character.")
		}
	}
}

func (S *Scanner) Identifier() {
	for S.IsAlphaNumber(S.Peek()) {
		S.Advance()
	}

    // other term sa = operator
	text := S.source[S.start:S.current]
	if text == "kay" {
		S.AddToken(types.EQUAL, nil)
		return
	}

    // other term sa == operator
	if text == "parehas" {
		S.AddToken(types.EQUAL_EQUAL, nil)
		return
	}

	t_type := keywords[text]
	if t_type == "" {
		t_type = types.IDETIFIER
	}

	S.AddToken(t_type, nil)
}

func (S *Scanner) Number() {
	for S.IsDigit(S.Peek()) {
		S.Advance()
	}

	if S.Peek() == '.' && S.IsDigit(S.PeekNext()) {
		S.Advance()

		for S.IsDigit(S.Peek()) {
			S.Advance()
		}
	}

	value, err := strconv.ParseFloat(S.source[S.start:S.current], 32)
	if err != nil {
		Error(S.line, "Cannot parse to Number")
	}

	S.AddToken(types.NUMBER, value)
}

func (S *Scanner) String() {
	for S.Peek() != '"' && S.IsAtEnd() == false {
		if S.Peek() == '\n' {
			S.line++
		}
		S.Advance()
	}

	if S.IsAtEnd() {
		Error(S.line, "Unterminated String")
		return
	}

	S.Advance()

	value := S.source[S.start+1 : S.current-1]
	S.AddToken(types.STRING, value)
}

func (S *Scanner) Match(expected byte) bool {
	if S.IsAtEnd() {
		return false
	}
	if S.source[S.current] != expected {
		return false
	}

	S.current++
	return true
}

func (S *Scanner) Peek() byte {
	if S.IsAtEnd() {
		return 0
	}
	return S.source[S.current]
}

func (S *Scanner) PeekNext() byte {
	if S.current+1 > len(S.source) {
		return 0
	}
	return S.source[S.current+1]
}

func (S *Scanner) IsAlpha(c byte) bool {
	var a byte = 97
	var z byte = 122
	var A byte = 65
	var Z byte = 90
	var underscore byte = 95

	return (c >= a && c <= z) ||
		(c >= A && c <= Z) ||
		(c == underscore)
}

func (S *Scanner) IsAlphaNumber(c byte) bool {
	return S.IsAlpha(c) || S.IsDigit(c)
}

func (S *Scanner) IsDigit(c byte) bool {
	var zero byte = 48
	var nine byte = 57

	return c >= zero && c <= nine
}

func (S *Scanner) Advance() byte {
	S.current++
	return S.source[S.current-1]
}

func (S *Scanner) AddToken(t_type types.TokenType, literal any) {
	text := S.source[S.start:S.current]
	S.tokens = append(S.tokens, *types.CreateToken(t_type, text, literal, S.line))
}

func (S *Scanner) IsAtEnd() bool {
	res := S.current >= len(S.source)
	return res
}
