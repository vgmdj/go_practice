package Longest_Common_Prefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefix(t *testing.T) {
	ast := assert.New(t)
	ast.Equal(longestCommonPrefix([]string{"abcdefg", "abc1123", "abdcefg"}), "ab")
	ast.Equal(longestCommonPrefix([]string{"aa", "a"}), "a")
	ast.Equal(longestCommonPrefix([]string{"a"}), "a")
}
