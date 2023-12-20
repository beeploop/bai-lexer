package parser

import (
	"github.com/BeepLoop/bai-interpreter/types"
)

func CallFactor() Expr {
	expr := CallUnary()

	if match(types.SLASH, types.STAR) {
		operator := previous()
		right := CallUnary()

		return CreateBinary(expr, operator, right)
	}

	return expr
}
