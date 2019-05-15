package leetcode

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestProblem20(t *testing.T) {
    assert.Equal(t, false, isValid("([)]"))
    assert.Equal(t, true, isValid("{}"))
    assert.Equal(t, true, isValid("()[]{}"))
    assert.Equal(t, false, isValid("(]"))
    assert.Equal(t, false, isValid("([)]"))
    assert.Equal(t, true, isValid("{()}"))
}


