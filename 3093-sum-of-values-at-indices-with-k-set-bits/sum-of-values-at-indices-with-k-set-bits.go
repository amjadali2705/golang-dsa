func sumIndicesWithKSetBits(nums []int, k int) int {
    sum := 0

    for i := 0; i < len(nums); i++ {
        if bits.OnesCount(uint(i)) == k {
            sum += nums[i]
        }
    }

    return sum
}