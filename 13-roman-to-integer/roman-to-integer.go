func romanToInt(s string) int {
    values := map[string]int{
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
    }

    sum := 0

    for i:=0; i<len(s); i++ {
        if i+1 < len(s) && values[string(s[i])] < values[string(s[i+1])] {
            sum -= values[string(s[i])]
        } else {
            sum += values[string(s[i])]
        }
    }
    return sum
}