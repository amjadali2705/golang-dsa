func isPalindrome(x int) bool {
    temp := x
    rev := 0
    
    for x > 0 {
        rem := x%10
        rev = rev*10+rem
        x /= 10
    }

    if rev == temp {
        return true
    } else {
        return false
    }
}