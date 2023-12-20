package parser

import "github.com/BeepLoop/bai-interpreter/types"

type Binary struct {
	Left     Expr
	Operator types.Token
	Right    Expr
}

func CreateBinary(left Expr, operator types.Token, right Expr) Binary {
	return Binary{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}
