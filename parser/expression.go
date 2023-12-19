package parser

import (
	"fmt"

	"github.com/BeepLoop/bai-interpreter/types"
)

type Expr interface{}

func Expression(token types.Token) Expr {

	fmt.Println(token)

	_ = CallEquality()

	return nil
}
