package leetcode

var dx = []int{-1, 0, 1, 0}
var dy = []int{0, 1, 0, -1}

func isRobotBounded(instructions string) bool {
    d := 0 
    x, y := 0, 0
    for _, r := range instructions {
        switch (r) {
        case 'G':
            x += dx[d]
            y += dy[d]
        case 'L':
            d = (d + 3) % 4
        case 'R':
            d = (d + 1) % 4
        }
    }
    return (x == 0 && y == 0) || (d != 0)
}
