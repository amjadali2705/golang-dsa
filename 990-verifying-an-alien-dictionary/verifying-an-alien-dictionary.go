func isAlienSorted(words []string, order string) bool {
    
    ranks := make([]int, 26)
    for i, char := range order {
        ranks[char - 'a'] = i
    }

    for i := 0; i < len(words)-1; i++ {
        w1 := words[i]
        w2 := words[i+1]
        
        isSorted := false

        for k := 0; k < len(w1) && k < len(w2); k++ {
            rank1 := ranks[w1[k] - 'a']
            rank2 := ranks[w2[k] - 'a']

            if rank1 < rank2 {
                isSorted = true
                break 
            }

            if rank1 > rank2 {
                return false
            }
        }

        if !isSorted {
            if len(w1) > len(w2) {
                return false
            }
        }
    }

    return true
}