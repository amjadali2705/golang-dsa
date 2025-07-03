func runningSum(nums []int) []int {
     runningSum := make([]int, len(nums))

    if len(nums) == 0 {
        return runningSum
    }

    runningSum[0] = nums[0]

    for i := 1; i < len(nums); i++ {
        runningSum[i] = runningSum[i-1] + nums[i]
    }

    return runningSum
}