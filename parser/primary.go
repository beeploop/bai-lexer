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

	// panic("Expect expression.")
	return errors.New("Expected expression")
}
