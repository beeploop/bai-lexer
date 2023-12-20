package parser

type Literal struct {
	Value any
}

func CreateLiteral(value any) Literal {
	return Literal{
		Value: value,
	}
}
