package main

func countWords(word string) int {
	count := 0

	for i := 0; i < len(word)-1; i++ {
		count++
	}
	return count
}
