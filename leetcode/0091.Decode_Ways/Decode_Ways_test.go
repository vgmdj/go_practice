package _091_Decode_Ways

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeWays(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(0, numDecodings("0"))
	ast.Equal(0, numDecodings("01"))
	ast.Equal(2, numDecodings("11"))
	ast.Equal(1, numDecodings("101"))
	ast.Equal(0, numDecodings("100"))
	ast.Equal(1, numDecodings("20"))
	ast.Equal(0, numDecodings("1324056"))
	ast.Equal(2, numDecodings("12"))
	ast.Equal(3, numDecodings("123"))
	ast.Equal(3, numDecodings("226"))
	ast.Equal(2, numDecodings("227"))
	ast.Equal(5, numDecodings("2226"))

}
