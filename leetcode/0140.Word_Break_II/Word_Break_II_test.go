package Word_Break_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBreak(t *testing.T) {
	ast := assert.New(t)

	ast.Equal([]string{"pine apple pen apple", "pine applepen apple", "pineapple pen apple"},
		wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))

}
