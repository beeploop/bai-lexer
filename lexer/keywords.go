package lexer

import "github.com/BeepLoop/bai-interpreter/types"

func Keywords() map[string]types.TokenType {
	keywords := make(map[string]types.TokenType)
	keywords["tapos"] = types.AND
	keywords["ay"] = types.ELSE
	keywords["atik"] = types.FALSE
	keywords["for"] = types.FOR
	keywords["fun"] = types.FUN
	keywords["ug"] = types.IF
	keywords["wala"] = types.NIL
	keywords["or"] = types.OR
	keywords["print"] = types.PRINT
	keywords["return"] = types.RETURN
	keywords["tinood"] = types.TRUE
	keywords["ang"] = types.VAR
	keywords["samtang"] = types.WHILE

	return keywords
}
