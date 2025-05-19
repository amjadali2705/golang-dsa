func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    if len(strs) == 1 {
        return strs[0]
    }

    sort.Strings(strs)

    for i := 0; i < len(strs[0]); i++ {
        if i >= len(strs[len(strs)-1]) || strs[0][i] != strs[len(strs)-1][i] {
            return strs[0][:i]
        }
    }
    return strs[0]
}