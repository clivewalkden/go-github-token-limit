package utils

import "strings"

func CenterString(str string, width int) string {
	spaces := int(float64(width-len(str)) / 2)
	// If the string is longer than the width, return the string as is
	if spaces < 0 {
		return str
	}
	return strings.Repeat(" ", spaces) + str + strings.Repeat(" ", width-(spaces+len(str)))
}

// ObscureToken Obscure a token by replacing the middle characters with "..."
func ObscureToken(token string) string {
	if len(token) < 8 {
		return token
	}
	return token[:14] + "..." + token[len(token)-4:]
}
