func generate(numRows int) [][]int {
	pascalTriangle := [][]int{}

	if numRows == 0 {
		return pascalTriangle
	}

	pascalTriangle = append(pascalTriangle, []int{1})

	for i := 1; i < numRows; i++ {
		prevRow := pascalTriangle[i-1]
		currRow := []int{1}

		for j := 1; j < len(prevRow); j++ {
			currRow = append(currRow, prevRow[j]+prevRow[j-1])
		}
		currRow = append(currRow, 1)
		pascalTriangle = append(pascalTriangle, currRow)
	}

	return pascalTriangle
}