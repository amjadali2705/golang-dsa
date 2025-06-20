func arrayPairSum(nums []int) int {
	sort.Ints(nums)

	maximizedSum := 0

	for i := 0; i < len(nums); i += 2 {
		maximizedSum += nums[i]
	}

	return maximizedSum
}