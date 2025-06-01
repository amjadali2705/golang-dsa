func thirdMax(nums []int) int {
    uniqueElemMap := make(map[int]bool)
    for _, num := range nums {
        uniqueElemMap[num] = true
    }

    uniqueElemSlice := []int{}
    for num := range uniqueElemMap {
        uniqueElemSlice = append(uniqueElemSlice, num)
    }

    sort.Slice(uniqueElemSlice, func(i, j int) bool {
        return uniqueElemSlice[i] > uniqueElemSlice[j]
    })

    if len(uniqueElemSlice) >= 3 {
        return uniqueElemSlice[2]
    } else {
        return uniqueElemSlice[0]
    }
}