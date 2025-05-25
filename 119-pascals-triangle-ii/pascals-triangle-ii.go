func getRow(rowIndex int) []int {
    if rowIndex == 0 {
        return []int{1}
    }

    currRow := []int{1}

    for i := 1; i <= rowIndex; i++ {
        newRow := make([]int, i+1)
        newRow[0] = 1

        for j := 1; j < i; j++ {
            newRow[j] = currRow[j-1] + currRow[j]
        } 
        newRow[i] = 1

        currRow = newRow
    }

    return currRow
}