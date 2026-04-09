package main

import (
	"fmt"
	"unicode"
)

func Token2(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		var classification string

		if unicode.IsLetter(r) {
			classification = "letter"
		} else if unicode.IsSpace(r) {
			classification = "whitespace"
		} else if r == '.' || r == ',' || r == ';' || r == ':' || r == '!' || r == '-' || r == '?' {
			classification = "punctuation"
		} else {
			classification = "other"
		}

		fmt.Println(i, string(r), classification)
	}
	return ""
}
