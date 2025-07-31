func maximizeSum(nums []int, k int) int {
    sort.Ints(nums)

	maxNum := nums[len(nums)-1]

	score := 0
	for i := 0; i < k; i++ {
		score += maxNum
		maxNum++
	}

	return score
}