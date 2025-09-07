func largeGroupPositions(s string) [][]int {
    result := [][]int{}

	start := 0
	l := 0 // length
	for i := range len(s) {
		if s[i] == s[start] {
			l++
		} else {
			if l >= 3 {
				result = append(result, []int{start, start + l - 1})
			}
            
			start = i
			l = 1
		}
	}

	if l >= 3 {
		result = append(result, []int{start, start + l - 1})
	}

	return result
}