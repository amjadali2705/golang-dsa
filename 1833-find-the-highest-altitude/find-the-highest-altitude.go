func largestAltitude(gain []int) int {
    highestAltitude := 0 
    currentAltitude := 0

    for _, g := range gain {
        currentAltitude += g 
        
        if currentAltitude > highestAltitude {
            highestAltitude = currentAltitude
        }
    }

    return highestAltitude
}