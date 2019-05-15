package leetcode

var numbers = map[rune]int{
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000,
}

func romanToInt(s string) int {
    res := 0
    last := -1
    for _, c := range s {
        v := numbers[c]
        if last != -1 && last < v {
            res -= last * 2
        }
        res += v
        last = v
    }
    return res
}
