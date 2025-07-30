func leftRightDifference(nums []int) []int {
    leftSum := make([]int, len(nums))
    rightSum := make([]int, len(nums))
    result := make([]int, len(nums))

    leftSum[0] = 0
    for i := 1; i < len(nums); i++ {
        leftSum[i] = leftSum[i-1] + nums[i-1]
    }

    rightSum[len(nums)-1] = 0
    for i := len(nums)-2; i >= 0; i-- {
        rightSum[i] = rightSum[i+1] + nums[i+1]
    }

    for i := 0; i < len(nums); i++ {
        result[i] = int(math.Abs(float64(leftSum[i]) - float64(rightSum[i])))
    }

    return result
}