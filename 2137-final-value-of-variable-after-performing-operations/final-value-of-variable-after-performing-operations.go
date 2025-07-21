func finalValueAfterOperations(operations []string) int {
    x := 0

    for _, operation := range operations {
        if operation == "++X" || operation == "X++" {
            x++
        }

        if operation == "--X" || operation == "X--" {
            x--
        }
    }

    return x
}