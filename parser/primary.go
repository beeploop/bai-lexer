package parser

import (
	"errors"
	"github.com/BeepLoop/bai-interpreter/types"
)

func CallPrimary() Expr {

	if match(types.FALSE) {
		return CreateLiteral(false)
	}

	if match(types.TRUE) {
		return CreateLiteral(true)
	}

	if match(types.NIL) {
		return CreateLiteral(nil)
	}

	if match(types.NUMBER, types.STRING) {
		token := previous().T_literal
		return CreateLiteral(token)
	}

	if match(types.LEFT_PAREN) {
		expr := Expression()
		consume(types.RIGHT_PAREN, "Expected ')' after expression.")
		return CreateGrouping(expr)
	}

	return errors.New("Expected expression")
}
