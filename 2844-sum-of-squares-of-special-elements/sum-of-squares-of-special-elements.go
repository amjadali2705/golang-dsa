func sumOfSquares(nums []int) int {
    sum := 0
    for i := 1; i <= len(nums); i++ {
        if len(nums) % i == 0 {
            sum += nums[i-1]*nums[i-1]
        }
    }
    return sum
}