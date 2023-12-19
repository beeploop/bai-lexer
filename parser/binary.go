package parser

import "github.com/BeepLoop/bai-interpreter/types"

type Binary struct {
	left     Expr
	operator types.Token
	right    Expr
}

func CreateBinary(left Expr, operator types.Token, right Expr) *Binary {
	return &Binary{
		left:     left,
		operator: operator,
		right:    right,
	}
}
