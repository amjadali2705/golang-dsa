func restoreString(s string, indices []int) string {
    shuffledStr := make([]rune, len(s))

    for i, char := range s {
        shuffledStr[indices[i]] = char
    }

    return string(shuffledStr)
}