package parser

type Expr interface{}

func Expression() Expr {
	return CallEquality()
}
