// https://leetcode.com/problems/last-stone-weight-ii/
package leetcode

func lastStoneWeightII(stones []int) int {
    MAX := 30 * 1000
    reach := make([]bool, MAX + 1)
    reach[0] = true

    sum := 0
    for _, stone := range stones {
        sum += stone
        for i := MAX - stone; i >= 0; i-- {
            if reach[i] {
                reach[i + stone] = true
            }
        }
    }

    for i := sum / 2; i >= 0; i-- {
        if reach[i] {
            return sum - i * 2
        }
    }

    return 0
}
