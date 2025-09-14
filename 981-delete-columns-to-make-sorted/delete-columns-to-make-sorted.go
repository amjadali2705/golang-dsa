func minDeletionSize(strs []string) int {
    if len(strs) == 0 {
		return 0
	}

	rows := len(strs)
	cols := len(strs[0])
	deletedCols := 0

	for i := 0; i < cols; i++ {
		for j := 0; j < rows-1; j++ {
			if strs[j][i] > strs[j+1][i] {
				deletedCols++
				break 
			}
		}
	}

	return deletedCols
}