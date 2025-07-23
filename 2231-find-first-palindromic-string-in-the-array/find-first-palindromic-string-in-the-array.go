func firstPalindrome(words []string) string {
    for _, word := range words {
        if isPalindrome(word) {
            return word
        }
    }
    return ""
}

func isPalindrome(word string) bool {
    start := 0
    end := len(word) - 1

    for start < end {
        if word[start] != word[end] {
            return false
        }
        start++
        end--
    }
    
    return true
}