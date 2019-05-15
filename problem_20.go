package leetcode

func isValid(s string) bool {
    var stack []rune
    for _, r := range s {
        switch r {
        case '(', '{', '[':
            stack = append(stack, r)
        case ')', '}', ']':
            if len(stack) == 0 {
                return false
            }
            last := stack[len(stack) - 1]
            if (r == '}' && last == '{') ||
                (r == ']' && last == '[') ||
                (r == ')' && last == '(') {
                stack = stack[:len(stack) - 1]
            } else {
                return false
            }
        }
    }
    return len(stack) == 0
}
