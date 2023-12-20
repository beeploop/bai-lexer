package parser

import (
	"reflect"

	"github.com/BeepLoop/bai-interpreter/types"
)

func match(tokens ...types.TokenType) bool {
	for _, token := range tokens {
		if checkType(token) {
			advance()
			return true
		}
	}

	return false
}

func isAtEnd() bool {
	if tokenList[curr].T_type == types.EOF {
		return true
	}

	return false
}

func checkType(token types.TokenType) bool {
	if isAtEnd() {
		return false
	}

	return peek().T_type == token
}

func peek() types.Token {
	return tokenList[curr]
}

func advance() types.Token {
	if isAtEnd() == false {
		curr++
	}

	return previous()
}

func previous() types.Token {
	if curr == 0 {
		return tokenList[curr]
	}

	return tokenList[curr-1]
}

func consume(token types.TokenType, message string) any {
	if checkType(token) {
		return advance()
	}

	return nil
}

func extractLiteralValue(expr Expr) float64 {
	val := reflect.ValueOf(expr).Elem()
	v := val.FieldByName("Value").Interface().(float64)

	return v
}
