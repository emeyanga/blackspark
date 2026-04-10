package tokenizer

import (
	"strings"
	"unicode"
)

func Tokenize(s string) string {
	runes := []rune(s)
	var buffer []rune
	var token []string

	for _, r := range runes {
		if unicode.IsLetter(r) {
			buffer = append(buffer, r)
		} else {
			if len(buffer) > 0 {
				token = append(token, string(buffer))
				buffer = buffer[:0]
			}
			if unicode.IsSpace(r) {
				continue
			} else {
				if r == '.' || r == ',' || r == ';' || r == ':' || r == '!' || r == '-' || r == '?' {
					token = append(token, string(r))
				}
			}
		}
	}
	if len(buffer) > 0 {
		token = append(token, string(buffer))
	}
	return JoinTokens(token)
}

func JoinTokens(token []string) string {
	return strings.Join(token, "")
}
