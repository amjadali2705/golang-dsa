func separateDigits(nums []int) []int {
    answer := []int{}

    for _, num := range nums {
        strNum := strconv.Itoa(num)
        for _, char := range strNum {
            digit, _ := strconv.Atoi(string(char))
            answer = append(answer, digit)
        }
    }

    return answer
}