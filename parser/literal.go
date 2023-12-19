package parser

type Literal struct {
	value any
}

func CreateLiteral(value any) *Literal {
	return &Literal{
		value: value,
	}
}
