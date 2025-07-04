func numIdenticalPairs(nums []int) int {
    freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	goodPairs := 0
	for _, count := range freqMap {
		if count >= 2 {
			goodPairs += (count * (count - 1)) / 2
		}
	}
	return goodPairs
}