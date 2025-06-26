func flipAndInvertImage(image [][]int) [][]int {
    n := len(image)

    for i := 0; i < n; i++ {
        // Step 1: Flip horizontally
        left := 0
        right := n - 1
        for left <= right {
            image[i][left], image[i][right] = image[i][right], image[i][left]
            left++
            right--
        }

        // Step 2: Invert
        for j := 0; j < n; j++ {
            if image[i][j] == 0 {
                image[i][j] = 1
            } else {
                image[i][j] = 0
            }
        }
    }

    return image
}