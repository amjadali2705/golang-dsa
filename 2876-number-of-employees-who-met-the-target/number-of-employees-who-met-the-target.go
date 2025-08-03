func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    count := 0
    for _, hour := range hours {
        if hour >= target {
            count++
        }
    }
    return count
}