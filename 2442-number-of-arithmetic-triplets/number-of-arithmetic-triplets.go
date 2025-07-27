func arithmeticTriplets(nums []int, diff int) int {
    numSet := make(map[int]bool)
    for _, num := range nums {
        numSet[num] = true
    }

    count := 0
    for _, num := range nums {
        if numSet[num + diff] && numSet[num + 2*diff] {
            count++
        }
    }

    return count
}