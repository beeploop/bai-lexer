package parser

import "github.com/BeepLoop/bai-interpreter/types"

type Unary struct {
	operator types.Token
	right    Expr
}

func CreateUnary(operator types.Token, right Expr) *Unary {
	return &Unary{
		operator: operator,
		right:    right,
	}
}

func CallUnary() Expr {

    return CallPrimary()
}
