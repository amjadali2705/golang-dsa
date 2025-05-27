func singleNumber(nums []int) int {
    res := 0

    for i, _ := range nums {
        res ^= nums[i]
    }

    return res
}