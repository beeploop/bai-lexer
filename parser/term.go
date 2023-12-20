package parser

import (
	"github.com/BeepLoop/bai-interpreter/types"
)

func CallTerm() Expr {
	expr := CallFactor()

	if match(types.MINUS, types.PLUS) {
		operator := previous()
		right := CallFactor()

		return CreateBinary(expr, operator, right)
	}

	return expr
}
