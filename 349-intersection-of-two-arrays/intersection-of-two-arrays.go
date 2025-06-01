func intersection(nums1 []int, nums2 []int) []int {
    set1 := make(map[int]bool)

    for _, num := range nums1 {
        set1[num] = true
    }

    var intersection []int

    for _, num := range nums2 {
        if _, present := set1[num]; present {
            intersection = append(intersection, num)
            delete(set1, num)
        }
    }

    return intersection
}