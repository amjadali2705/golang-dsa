func canConstruct(ransomNote string, magazine string) bool {
    charCounts := make([]int, 26)

    for _, char := range magazine {
        charCounts[char - 'a']++
    }

    for _, char := range ransomNote {
        if charCounts[char - 'a'] == 0 {
            return false
        }

        charCounts[char - 'a']--
    }

    return true
}