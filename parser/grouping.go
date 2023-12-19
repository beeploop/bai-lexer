package parser

type Grouping struct {
	expression Expr
}

func CreateGrouping(expression Expr) *Grouping {
	return &Grouping{
		expression: expression,
	}
}
