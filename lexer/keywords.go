package lexer

import "github.com/BeepLoop/bai-interpreter/types"

func Keywords() map[string]types.TokenType {
	keywords := make(map[string]types.TokenType)
	keywords["ang"] = types.VAR

	keywords["tapos"] = types.AND
	keywords["or"] = types.OR

	keywords["ug"] = types.IF
	keywords["ay-ug"] = types.ELIF
	keywords["edi"] = types.ELSE

	keywords["tinood"] = types.TRUE
	keywords["atik"] = types.FALSE

	keywords["kada"] = types.FOR
	keywords["samtang"] = types.WHILE

	keywords["katikaran"] = types.FUN
	keywords["payts"] = types.RETURN
	keywords["pakita"] = types.PRINT

	keywords["wala"] = types.NIL

	return keywords
}
