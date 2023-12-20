package parser

import (
	"github.com/BeepLoop/bai-interpreter/types"
)

func CallComparison() Expr {
	expr := CallTerm()

	if match(types.GREATER, types.GREATER_EQUAL, types.LESS, types.LESS_EQUAL) {
		operator := previous()
		right := CallTerm()

		return CreateBinary(expr, operator, right)
	}

	return expr
}
