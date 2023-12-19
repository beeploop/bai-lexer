package parser

import "github.com/BeepLoop/bai-interpreter/types"

type Binary struct {
	left     Expression
	operator types.Token
	right    Expression
}

func CreateBinary(left Expression, operator types.Token, right Expression) *Binary {
	return &Binary{
		left:     left,
		operator: operator,
		right:    right,
	}
}
