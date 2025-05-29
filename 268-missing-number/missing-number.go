func missingNumber(nums []int) int {
    n := len(nums)
    totalSum := n*(n+1)/2

    sum := 0
    for i, _ := range nums {
        sum += nums[i]
    }

    missingNo := totalSum - sum
    return missingNo
}