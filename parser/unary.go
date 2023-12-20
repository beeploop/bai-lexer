package parser

import (
	"github.com/BeepLoop/bai-interpreter/types"
)

type Unary struct {
	Operator types.Token
	Right    Expr
}

func CreateUnary(operator types.Token, right Expr) Unary {
	return Unary{
		Operator: operator,
		Right:    right,
	}
}

func CallUnary() Expr {

	if match(types.BANG, types.MINUS) {
		operator := previous()
		right := CallUnary()

		return CreateUnary(operator, right)
	}

	return CallPrimary()
}
