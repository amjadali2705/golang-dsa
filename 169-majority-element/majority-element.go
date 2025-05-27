func majorityElement(nums []int) int {
    counts := make(map[int]int)

    for _, num := range nums {
        counts[num]++
    }

    majorityThreshold := len(nums) / 2

    for num, count := range counts {
        if count > majorityThreshold {
            return num
        }
    }

    return -1
}