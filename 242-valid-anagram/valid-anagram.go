func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    charCounts := make(map[rune]int)

    for _, char := range s {
        charCounts[char]++
    }

    for _, char := range t {
        if count, ok := charCounts[char]; !ok || count == 0 {
            return false
        }
        charCounts[char]--
    }

    return true
}