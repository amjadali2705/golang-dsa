func findDisappearedNumbers(nums []int) []int {
    numSet := make(map[int]bool)
    for _, num := range nums {
        numSet[num] = true
    }

    missingNum := []int{}
    for i := 1; i <= len(nums); i++ {
        if _, found := numSet[i]; !found {
            missingNum = append(missingNum, i)
        }
    }

    return missingNum
}