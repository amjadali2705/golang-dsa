func uncommonFromSentences(s1 string, s2 string) []string {
    combined := s1 + " " + s2
	words := strings.Fields(combined)

	wordCounts := make(map[string]int)
	for _, word := range words {
		wordCounts[word]++
	}

	var uncommonWords []string
	for word, count := range wordCounts {
		if count == 1 {
			uncommonWords = append(uncommonWords, word)
		}
	}

	return uncommonWords
}