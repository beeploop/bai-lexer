package parser

type Grouping struct {
	expression Expression
}

func CreateGrouping(expression Expression) *Grouping {
	return &Grouping{
		expression: expression,
	}
}
