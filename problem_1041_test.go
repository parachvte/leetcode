package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func TestProblem1041(t *testing.T) {
    assert.Equal(t, isRobotBounded("GGLLGG"), true)
    assert.Equal(t, isRobotBounded("GG"), false)
    assert.Equal(t, isRobotBounded("GL"), true)
}

