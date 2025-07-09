func maximumWealth(accounts [][]int) int {
    maxWealth := 0

    for _, customerAccounts := range accounts {
        currentWealth := 0
        for _, amount := range customerAccounts {
            currentWealth += amount
        }

        if currentWealth > maxWealth {
            maxWealth = currentWealth
        }
    }

    return maxWealth
}