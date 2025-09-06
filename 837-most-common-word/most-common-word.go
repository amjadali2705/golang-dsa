func mostCommonWord(paragraph string, banned []string) string {
    bannedMap := make(map[string]struct{})
	for _, word := range banned {
		bannedMap[word] = struct{}{}
	}

	freqMap := make(map[string]int)
	replacer := strings.NewReplacer("!", " ", "?", " ", "'", " ", ",", " ", ";", " ", ".", " ")
	processedParagraph := replacer.Replace(paragraph)

	words := strings.Fields(strings.ToLower(processedParagraph))

	for _, word := range words {
		if _, ok := bannedMap[word]; !ok {
			freqMap[word]++
		}
	}

	maxCount := 0
	result := ""

	for word, count := range freqMap {
		if count > maxCount {
			maxCount = count
			result = word
		}
	}

	return result
}