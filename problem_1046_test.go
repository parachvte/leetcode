package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func TestProblem1046(t *testing.T) {
    assert.Equal(t, lastStoneWeight([]int{2, 7, 4, 1, 8, 1}), 1, "Not Equal")
}
