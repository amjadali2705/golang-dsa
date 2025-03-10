package main

func QuickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}

	pivotIndex := len(list) / 2
	pivotElem := list[pivotIndex]

	less := []int{}
	greater := []int{}

	for i, v := range list {
		if i != pivotIndex {
			if v < pivotElem {
				less = append(less, v)
			} else {
				greater = append(greater, v)
			}
		}
	}

	sortedLess := QuickSort(less)
	sortedGreater := QuickSort(greater)

	sorted := append(sortedLess, pivotElem)
	sorted = append(sorted, sortedGreater...)

	return sorted
}