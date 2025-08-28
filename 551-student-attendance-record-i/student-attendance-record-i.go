func checkRecord(s string) bool {
    absent, late := 0, 0
	for i := range s {
		if s[i] == 'A' {
			absent++

			if absent == 2 {
				return false
			}
		}

		if s[i] == 'L' {
			late++

			if late == 3 {
				return false
			}
		} else {
			late = 0
		}
	}
	return true
}