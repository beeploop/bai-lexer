package parser

import (
	"github.com/BeepLoop/bai-interpreter/types"
)

func CallEquality() Expr {
	expr := CallComparison()

	if match(types.BANG_EQUAL, types.EQUAL_EQUAL) {
		operator := previous()
		right := CallComparison()

		return CreateBinary(expr, operator, right)
	}

	return expr
}
