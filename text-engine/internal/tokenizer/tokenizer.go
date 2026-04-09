package tokenizer

import (
	"unicode"
)

func Tokenize(s string) []string {
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
				token = append(token, string(buffer))
			}
		}
	}
	if len(buffer) > 0 {
		token = append(token, string(buffer))
	}
	return token
}
