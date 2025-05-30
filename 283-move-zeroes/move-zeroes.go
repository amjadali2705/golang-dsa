func moveZeroes(nums []int)  {
    nonZeroPtr := 0
    for i, _ := range nums {
        if nums[i] != 0 {
            nums[nonZeroPtr] = nums[i]
            nonZeroPtr++
        }
    } 

    for i := nonZeroPtr; i < len(nums); i++ {
        nums[i] = 0
    }
}