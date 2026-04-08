package main

import (
	"fmt"
	"strings"
)

func reverseSlice(word string) string {
	strings.Fields(word)

	for i := 0; i < len(word)-1; i-- {
		return ""
	}
	return ""
}

func main() {
	fmt.Println(reverseSlice("excel"))
}
