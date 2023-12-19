package parser

import "github.com/BeepLoop/bai-interpreter/types"

type Unary struct {
	operator types.Token
	right    Expression
}

func CreateUnary(operator types.Token, right Expression) *Unary {
	return &Unary{
		operator: operator,
		right:    right,
	}
}
