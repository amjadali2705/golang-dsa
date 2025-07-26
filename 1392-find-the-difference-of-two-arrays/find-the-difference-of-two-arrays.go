func findDifference(nums1 []int, nums2 []int) [][]int {
    setNums1 := make(map[int]bool)
    for _, num := range nums1 {
        setNums1[num] = true
    }

    setNums2 := make(map[int]bool)
    for _, num := range nums2 {
        setNums2[num] = true
    }

    var diff1 []int
    for num, _ := range setNums1 {
        if _, found := setNums2[num]; !found {
            diff1 = append(diff1, num)
        }
    }

    var diff2 []int
    for num, _ := range setNums2 {
        if _, found := setNums1[num]; !found {
            diff2 = append(diff2, num)
        }
    }
    
    return [][]int{diff1, diff2}
}