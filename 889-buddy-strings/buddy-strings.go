func buddyStrings(s string, goal string) bool {
    if len(s) != len(goal) {
        return false
    }

    if s == goal {
        seen := make(map[rune]bool)
        for _, char := range s {
            if seen[char] {
                return true
            }

            seen[char] = true
        }
        
        return false
    }

    diff := []int{}
    for i := 0; i < len(s); i++ {
        if s[i] != goal[i] {
            diff = append(diff, i)
        }
    }

    if len(diff) == 2 {
        if s[diff[0]] == goal[diff[1]] && s[diff[1]] == goal[diff[0]] {
            return true
        }
    }

    return false
}