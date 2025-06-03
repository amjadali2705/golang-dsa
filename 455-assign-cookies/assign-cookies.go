func findContentChildren(g []int, s []int) int {
    i := 0
    j := 0
    res := 0

    sort.Ints(g)
    sort.Ints(s)

    for i < len(g) && j < len(s) {
        if s[j] >= g[i] {
            i++
            j++
            res++
        } else {
            j++
        }
    }

    return res
}