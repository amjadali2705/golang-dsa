func maximumNumberOfStringPairs(words []string) int {
    freqMap := make(map[string]int)
    count := 0

    for _, word := range words {
        reverseWord := reverseString(word)
        if freqMap[reverseWord] > 0 {
            count++
            freqMap[reverseWord]--
        } else {
            freqMap[word]++
        }
    }
    return count
}

func reverseString(word string) string {
    i := 0;
    j := len(word) - 1;

    runes := []rune(word)
    for i < j {
        runes[i], runes[j] = runes[j], runes[i]
        i++
        j--
    }
    return string(runes)
}