func differenceOfSum(nums []int) int {
    elemSum := 0
    digiSum := 0

    for _, num := range nums {
        elemSum += num
    }

    for _, num := range nums {
        currNum := num
        for currNum > 0 {
            rem := currNum % 10
            digiSum += rem
            currNum /= 10
        }
    }

    return int(math.Abs(float64(elemSum - digiSum)))
}