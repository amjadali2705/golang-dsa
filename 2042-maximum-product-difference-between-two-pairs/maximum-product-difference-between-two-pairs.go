func maxProductDifference(nums []int) int {
    // Approach 1: Sorting

    // sort.Ints(nums)

	// n := len(nums)
	// productMax := nums[n-1] * nums[n-2]
	// productMin := nums[0] * nums[1]
	// return productMax - productMin

    // Approach 2: without sorting
    max1, max2 := math.MinInt, math.MinInt
	min1, min2 := math.MaxInt, math.MaxInt

	for _, num := range nums {
		if num >= max1 { 
			max2 = max1
			max1 = num
		} else if num > max2 {
			max2 = num
		}

		if num <= min1 { 
			min2 = min1
			min1 = num
		} else if num < min2 {
			min2 = num
		}
	}

	return (max1 * max2) - (min1 * min2)
}