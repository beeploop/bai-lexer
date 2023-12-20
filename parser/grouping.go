package parser

type Grouping struct {
	Expression Expr
}

func CreateGrouping(expression Expr) Grouping {
	return Grouping{
		Expression: expression,
	}
}
