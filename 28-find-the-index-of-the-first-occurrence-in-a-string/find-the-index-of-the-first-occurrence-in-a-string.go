func strStr(haystack string, needle string) int {
    n := len(haystack)
    m := len(needle)

    if m > n {
        return -1
    }

    for i := 0; i <= n-m; i++ {
        if haystack[i:i+m] == needle {
            return i 
        }
    }
    
    return -1

}