package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func TestProblem13(t *testing.T) {
    assert.Equal(t, romanToInt("III"),  3, "Not Equal")
    assert.Equal(t, romanToInt("IV"), 4) 
    assert.Equal(t, romanToInt("IX"), 9)
    assert.Equal(t, romanToInt("LVIII"), 58)
    assert.Equal(t, romanToInt("MCMXCIV"), 1994)
}
