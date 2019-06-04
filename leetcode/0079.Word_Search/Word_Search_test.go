package Word_Search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	ast := assert.New(t)

	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	ast.Equal(true, exist(board, "ABCCED"))

	ast.Equal(true, exist(board, "SEE"))

	ast.Equal(false, exist(board, "ABCB"))

}
