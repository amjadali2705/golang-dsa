func findJudge(n int, trust [][]int) int {
    if n == 1 {
        return 1
    }

    score := make([]int, n+1)

    for _, relation := range trust {
        a := relation[0] 
        b := relation[1]

        score[a]--
        score[b]++
    }

    judgeScore := n - 1
    
    for i := 1; i <= n; i++ {
        if score[i] == judgeScore {
            return i
        }
    }

    return -1
}