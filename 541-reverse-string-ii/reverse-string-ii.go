func reverseStr(s string, k int) string {
    str := []byte(s)
	n := len(str)

	for i := 0; i < n; i += 2 * k {
		
		end := i + k
		if end > n {
			end = n
		}
		reverse(str[i:end])
	}

	return string(str)
}

func reverse(b []byte) {
	l, r := 0, len(b)-1
	for l < r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
}