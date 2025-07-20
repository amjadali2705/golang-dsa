func countKDifference(nums []int, k int) int {
    count := 0
    freqMap := make(map[int]int)

    for _, num := range nums {
        if val, ok := freqMap[num-k]; ok {
            count += val
        }

        if val, ok := freqMap[num+k]; ok {
            count += val
        }
        freqMap[num]++
    }

    return count
}