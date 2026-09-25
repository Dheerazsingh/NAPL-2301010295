package strop

import "strings"

func CountVowels(s string) int {
	count := 0
	vowels := "aeiouAEIOU"
	for _, ch := range s {
		if strings.ContainsRune(vowels, ch) {
			count++
		}
	}
	return count
}
