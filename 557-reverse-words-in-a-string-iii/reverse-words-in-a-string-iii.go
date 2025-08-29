func reverseWords(s string) string {
    words := strings.Fields(s) 
    for i, word := range words {
        reversedWord := reverseString(word)
        words[i] = reversedWord
    }
    
    return strings.Join(words, " ")
}

func reverseString(s string) string {
    runes := []rune(s)
    left, right := 0, len(runes)-1
    for left < right {
        runes[left], runes[right] = runes[right], runes[left]
        left++
        right--
    }

    return string(runes)
}