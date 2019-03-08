package Convert_a_Number_to_Hexadecimal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToHex(t *testing.T) {
	ast := assert.New(t)

	ast.Equal("0",toHex(0))
	ast.Equal("1", toHex(1))
	ast.Equal("a", toHex(10))
	ast.Equal("f", toHex(15))
	ast.Equal("10", toHex(16))
	ast.Equal("11", toHex(17))
	ast.Equal("159", toHex(345))
	ast.Equal("234", toHex(564))
	ast.Equal("ffffffff",toHex(-1))

}
