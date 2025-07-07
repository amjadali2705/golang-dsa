func sumOddLengthSubarrays(arr []int) int {
    n := len(arr)
    sum := 0

    prefSum := make([]int, n)

    prefSum[0] = arr[0]
    for i := 1; i < n; i++ {
        prefSum[i] = prefSum[i-1] + arr[i]
    }

    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            currentLength := j - i + 1

            if currentLength % 2 == 1 {
                if i == 0 {
                    sum += prefSum[j]
                } else {
                    sum += prefSum[j] - prefSum[i-1]
                }
            }
        }
    }
    return sum
}