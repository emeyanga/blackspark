package main

import (
	"fmt"
)

func token(s string) string {
	runes := []rune(s)
	// Reasoning: Should we range over string directly or convert to []rune first? Why?  i'll convert to rune becase that way it covers unicode characters, andis easily convertible to bytes
	for i := range runes {
		fmt.Println(i, string(runes[i]))
	}
	return ""
}

func main() {
	fmt.Println(token("Hello world. How are you?"))
}
