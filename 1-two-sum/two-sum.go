func twoSum(nums []int, target int) []int {
    hashMap := make(map[int]int)
    for i, num := range nums {
        complement := target - num

        j, ok := hashMap[complement]
        if ok {
            return []int{j, i}
        }
        
        hashMap[num] = i
    }

    return nil  
}