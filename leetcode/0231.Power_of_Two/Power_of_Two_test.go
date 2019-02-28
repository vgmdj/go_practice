package Power_of_Two

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPower(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(true, isPowerOfTwo(1))
	ast.Equal(true, isPowerOfTwo(2))
	ast.Equal(false, isPowerOfTwo(3))
	ast.Equal(true, isPowerOfTwo(4))
	ast.Equal(false, isPowerOfTwo(5))

}
