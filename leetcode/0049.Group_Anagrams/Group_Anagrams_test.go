package Group_Anagrams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	ast := assert.New(t)

	ast.Equal([][]string{
		[]string{"eat", "tea", "ate"},
		[]string{"tan", "nat"},
		[]string{"bat"},
	},
		groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))

}
