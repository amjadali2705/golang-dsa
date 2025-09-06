func shortestToChar(s string, c byte) []int {
    n := len(s)
	answer := make([]int, n)

	prev := -n
	for i := 0; i < n; i++ {
		if s[i] == c {
			prev = i
		}
		answer[i] = i - prev
	}

	prev = 2 * n
	for i := n - 1; i >= 0; i-- {
		if s[i] == c {
			prev = i
		}
		answer[i] = int(math.Min(float64(answer[i]), float64(prev-i)))
	}

	return answer
}