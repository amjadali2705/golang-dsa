func createTargetArray(nums []int, index []int) []int {
    target := []int{} 

	for i := 0; i < len(nums); i++ { 
		val := nums[i]  
		idx := index[i] 

		target = append(target[:idx], append([]int{val}, target[idx:]...)...)
	}

	return target
}