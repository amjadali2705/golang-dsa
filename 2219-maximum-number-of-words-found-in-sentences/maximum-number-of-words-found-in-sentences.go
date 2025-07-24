func mostWordsFound(sentences []string) int {
    maxWords := 0 

    for _, sentence := range sentences {
        words := strings.Fields(sentence)
        currentWordsCount := len(words)

        if currentWordsCount > maxWords {
            maxWords = currentWordsCount
        }
    }

    return maxWords
}