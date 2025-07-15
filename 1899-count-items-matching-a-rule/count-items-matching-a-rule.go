func countMatches(items [][]string, ruleKey string, ruleValue string) int {
    matchCount := 0
    var ruleIndex int

    switch ruleKey {
    case "type":
        ruleIndex = 0
    case "color":
        ruleIndex = 1
    case "name":
        ruleIndex = 2
    }

    for _, item := range items {
        if item[ruleIndex] == ruleValue {
            matchCount++
        }
    }

    return matchCount
}