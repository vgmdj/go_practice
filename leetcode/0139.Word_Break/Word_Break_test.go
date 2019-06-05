package Word_Break

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordBreak(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(false, wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
		[]string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}))
	ast.Equal(true, wordBreak("leetcode", []string{"leet", "code"}))
	ast.Equal(false, wordBreak("appsletter", []string{"app", "letter"}))
	ast.Equal(true, wordBreak("appletter", []string{"app", "letter"}))

}
