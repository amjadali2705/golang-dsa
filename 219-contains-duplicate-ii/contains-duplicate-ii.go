func containsNearbyDuplicate(nums []int, k int) bool {
    lastIndx := make(map[int]int)

    for i, num := range nums {
        prevIndx, found := lastIndx[num]
        if found && i - prevIndx <= k {
            return true
        }

        lastIndx[num] = i
    }

    return false
}