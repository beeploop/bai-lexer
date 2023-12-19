package lexer

type Expression interface{}

type Binary struct {
	left     Expression
	operator Token
	right    Expression
}

func CreateBinary(left Expression, operator Token, right Expression) *Binary {
	return &Binary{
		left:     left,
		operator: operator,
		right:    right,
	}
}

type Grouping struct {
	expression Expression
}

func CreateGrouping(expression Expression) *Grouping {
	return &Grouping{
		expression: expression,
	}
}

type Literal struct {
	value any
}

func CreateLiteral(value any) *Literal {
	return &Literal{
		value: value,
	}
}

type Unary struct {
	operator Token
	right    Expression
}

func CreateUnary(operator Token, right Expression) *Unary {
	return &Unary{
		operator: operator,
		right:    right,
	}
}

