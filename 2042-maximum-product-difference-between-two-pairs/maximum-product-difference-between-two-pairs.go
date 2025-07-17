func maxProductDifference(nums []int) int {
    sort.Ints(nums)

	n := len(nums)

	productMax := nums[n-1] * nums[n-2]

	productMin := nums[0] * nums[1]

	return productMax - productMin
}