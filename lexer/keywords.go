package lexer

func Keywords() map[string]TokenType {
	keywords := make(map[string]TokenType)
	keywords["tapos"] = AND
	keywords["class"] = CLASS
	keywords["ay"] = ELSE
	keywords["atik"] = FALSE
	keywords["for"] = FOR
	keywords["fun"] = FUN
	keywords["ug"] = IF
	keywords["nil"] = NIL
	keywords["or"] = OR
	keywords["print"] = PRINT
	keywords["return"] = RETURN
	keywords["super"] = SUPER
	keywords["kini"] = THIS
	keywords["tinood"] = TRUE
	keywords["ang"] = VAR
	keywords["samtang"] = WHILE

	return keywords
}
