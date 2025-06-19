func findPoisonedDuration(timeSeries []int, duration int) int {
    totalPoisonedTime := 0

    for i := 0; i < len(timeSeries) - 1; i++ {
        currentPoisonEnd := timeSeries[i] + duration

        if currentPoisonEnd <= timeSeries[i+1] {
            totalPoisonedTime += duration
        } else {
            totalPoisonedTime += timeSeries[i+1] - timeSeries[i]
        }
    }

    totalPoisonedTime += duration

    return totalPoisonedTime
}