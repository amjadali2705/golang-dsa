func intersect(nums1 []int, nums2 []int) []int {
    freqMap := make(map[int]int) 

    for _, num := range nums1 {
        freqMap[num]++
    }

    var intersection []int

    for _, num := range nums2 {
        if count, present := freqMap[num]; present && count > 0 {
            intersection = append(intersection, num)
            freqMap[num]--
        }
    }

    return intersection
}