func numberOfLines(widths []int, s string) []int {
    width, line := 0, 1
    for i := range s {
        width += widths[s[i] - 'a']
        if width > 100 {
            line++
            width = widths[s[i] - 'a']
        }
    }

    return []int{line, width}
}