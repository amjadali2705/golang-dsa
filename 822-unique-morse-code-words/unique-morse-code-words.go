func uniqueMorseRepresentations(words []string) int {
    morseMap := map[byte]string {
        'a': ".-",
        'b': "-...",
        'c': "-.-.",
        'd': "-..",
        'e': ".",
        'f': "..-.",
        'g': "--.",
        'h': "....",
        'i': "..",
        'j': ".---",
        'k': "-.-",
        'l': ".-..",
        'm': "--",
        'n': "-.",
        'o': "---",
        'p': ".--.",
        'q': "--.-",
        'r': ".-.",
        's': "...",
        't': "-",
        'u': "..-",
        'v': "...-",
        'w': ".--",
        'x': "-..-",
        'y': "-.--",
        'z': "--..",
    }

    morseSet := make(map[string]bool)

    for _, word := range words {
        morseWord := ""

        for _, letter := range []byte(word) {
            morseWord += morseMap[letter]
        }

        morseSet[morseWord] = true
    }

    return len(morseSet)
}