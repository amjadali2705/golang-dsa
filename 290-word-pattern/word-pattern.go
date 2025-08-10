func wordPattern(pattern string, s string) bool {
    words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	patternToWord := make(map[rune]string)
	wordToPattern := make(map[string]rune)

	for i, char := range pattern {
		word := words[i]

		if mappedWord, ok := patternToWord[char]; ok {
			if mappedWord != word {
				return false
			}
		} else {
			if mappedChar, ok := wordToPattern[word]; ok {
				if mappedChar != char {
					return false
				}
			}
            
			patternToWord[char] = word
			wordToPattern[word] = char
		}
	}

	return true
}