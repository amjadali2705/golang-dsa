func sortPeople(names []string, heights []int) []string {
    heightToName := make(map[int]string)
	for i := 0; i < len(names); i++ {
		heightToName[heights[i]] = names[i]
	}

	var sortedHeights []int
	for height := range heightToName {
		sortedHeights = append(sortedHeights, height)
	}

	
	sort.Slice(sortedHeights, func(i, j int) bool {
		return sortedHeights[i] > sortedHeights[j] 
	})

	resultNames := make([]string, len(sortedHeights))
	for i, height := range sortedHeights {
		resultNames[i] = heightToName[height]
	}

	return resultNames
}