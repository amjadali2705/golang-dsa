func findMaxConsecutiveOnes(nums []int) int {
    count := 0
    maxCount := 0

    for _, num := range nums {
        if num == 1 {
            count++
        } else {
            if count > maxCount {
                maxCount = count
            }
            count = 0
        }
    }

    if count > maxCount {
        maxCount = count
    }

    return maxCount
}