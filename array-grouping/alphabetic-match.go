package main

import (
	"fmt"
)

func main() {
	words := []string{"cook", "save", "taste", "vase", "aves", "state", "map"}
	groupedWords := proceedWords(words)
	fmt.Println(groupedWords)
}

// proceedWords groups words that contain every character of another word.
func proceedWords(words []string) [][]string {
	var groupedWords [][]string

	for i := 0; i < len(words); i++ {
		match := false
		for j := 0; j < len(groupedWords); j++ {
			if containsEveryChar(groupedWords[j][0], words[i]) {
				groupedWords[j] = append(groupedWords[j], words[i])
				match = true
				break
			}
		}
		if !match {
			groupedWords = append(groupedWords, []string{words[i]})
		}
	}

	return groupedWords
}

// containsEveryChar checks if one string contains every character of another string.
func containsEveryChar(srcStr, dstStr string) bool {
	charCounts := make(map[rune]int)

	for _, ch := range dstStr {
		if _, ok := charCounts[ch]; !ok {
			charCounts[ch] = 1
		} else {
			charCounts[ch]++
		}
	}

	for _, ch := range srcStr {
		if count, ok := charCounts[ch]; ok {
			charCounts[ch] = count - 1
			if charCounts[ch] == 0 {
				delete(charCounts, ch)
			}
		}
	}

	return len(charCounts) == 0
}
