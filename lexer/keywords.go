package lexer

func Keywords() map[string]TokenType {
	keywords := make(map[string]TokenType)
	keywords["tapos"] = AND
	keywords["ay"] = ELSE
	keywords["atik"] = FALSE
	keywords["for"] = FOR
	keywords["fun"] = FUN
	keywords["ug"] = IF
	keywords["wala"] = NIL
	keywords["or"] = OR
	keywords["print"] = PRINT
	keywords["return"] = RETURN
	keywords["tinood"] = TRUE
	keywords["ang"] = VAR
	keywords["samtang"] = WHILE

	return keywords
}
