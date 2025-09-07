func toGoatLatin(sentence string) string {
    words := strings.Fields(sentence)
    vowels := "aeiouAEIOU"

    for i, word := range words {
        if strings.ContainsRune(vowels, rune(word[0])) {
            words[i] = word + "ma"
        } else {
            words[i] = word[1:] + string(word[0]) + "ma"
        }

        words[i] = words[i] + strings.Repeat("a", i+1)
    }

    return strings.Join(words, " ")
}