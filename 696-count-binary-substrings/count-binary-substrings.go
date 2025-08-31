func countBinarySubstrings(s string) int {
    ans := 0
    prevCount := 0
    currCount := 1

    for i := 1; i < len(s); i++ {
        if s[i] == s[i-1] {
            currCount++
        } else {
            ans += min(prevCount, currCount)
            prevCount = currCount
            currCount = 1
        }
    }
    
    ans += min(prevCount, currCount)
    
    return ans
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}