func detectCapitalUse(word string) bool {
    capitalCount := 0
    for _, r := range word {
        if unicode.IsUpper(r) {
            capitalCount++
        }
    }

    if capitalCount == len(word) {
        return true
    }

    if capitalCount == 0 {
        return true
    }

    if capitalCount == 1 && unicode.IsUpper(rune(word[0])) {
        return true
    }

    return false
}