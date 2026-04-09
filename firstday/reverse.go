package main

func reverseSlice(word string) string {
	result := ""

	for i := len(word) - 1; i >= 0; i-- {
		result += string(word[i])
	}
	return result
}

func reverseStrings(word string) string {
	runes := []rune(word)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
