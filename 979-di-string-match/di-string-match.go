func diStringMatch(s string) []int {
    n := len(s)
	perm := make([]int, n+1)
	low, high := 0, n

	for i := 0; i < n; i++ {
		if s[i] == 'I' {
			perm[i] = low
			low++
		} else { 
			perm[i] = high
			high--
		}
	}
    
	perm[n] = low

	return perm
}