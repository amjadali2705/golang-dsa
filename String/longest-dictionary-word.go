package main

import "strings"

func LongestDictionaryWordContainingKey(key string, dic []string) string {
	longestWord := ""

	for _, word := range dic {
		if containsAllLetters(word, key) && len(word) > len(longestWord) {
			longestWord = word
		}
	}
	return longestWord
}

func containsAllLetters(word, key string) bool {
	for _, char := range key {
		if !strings.ContainsRune(word, char) {
			return false
		}
	}
	return true
}
