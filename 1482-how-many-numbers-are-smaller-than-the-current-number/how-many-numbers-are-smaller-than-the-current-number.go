func smallerNumbersThanCurrent(nums []int) []int {
    n := len(nums)
    result := make([]int, n)

    for i := 0; i < n; i++ {
        count := 0 
        for j := 0; j < n; j++ {
            if j != i && nums[j] < nums[i] {
                count++
            }
        }
        result[i] = count
    }

    return result
}