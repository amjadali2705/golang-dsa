func sumCounts(nums []int) int {
    totalSumOfSquares := 0

	for i := 0; i < len(nums); i++ {
		distinctElements := make(map[int]bool)
		for j := i; j < len(nums); j++ {
			distinctElements[nums[j]] = true
			distinctCount := len(distinctElements)
			totalSumOfSquares += distinctCount * distinctCount
		}
	}

	return totalSumOfSquares
}