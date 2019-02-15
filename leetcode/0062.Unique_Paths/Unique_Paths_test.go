package Unique_Paths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePath(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(3, uniquePaths(3, 2))
	ast.Equal(28, uniquePaths(7, 3))
	ast.Equal(1101716330, uniquePaths(38, 10))

}
