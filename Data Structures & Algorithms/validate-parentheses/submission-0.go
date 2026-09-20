func isValid(s string) bool {
    stack := []rune{}
    pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}

    for _, c := range s {
        if open, exists := pairs[c]; exists {
            if len(stack) > 0 {
                top := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                if top != open {
                    return false
                }
            } else {
                return false
            }
        } else {
            stack = append(stack, c)
        }
    }

    return len(stack) == 0
}