func findTheDifference(s string, t string) byte {
    sumS, sumT := 0, 0

    for _, char := range s {
        sumS += int(char)
    }

    for _, char := range t {
        sumT += int(char)
    }

    diff := sumT - sumS

    return byte(diff)
}