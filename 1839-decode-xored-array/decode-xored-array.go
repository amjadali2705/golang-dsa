func decode(encoded []int, first int) []int {
    n := len(encoded) + 1
    arr := make([]int, n)

    arr[0] = first

    for i := 0; i < n-1; i++ {
        arr[i+1] = arr[i] ^ encoded[i]
    }

    return arr
}