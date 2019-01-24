package ZigZag_Conversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZigZag(t *testing.T) {
	ast := assert.New(t)

	ast.Equal("acebdf", convert("abcdef", 2))
	ast.Equal("PAHNAPLSIIGYIR", convert("PAYPALISHIRING", 3))
	ast.Equal("PINALSIGYAHRPI", convert("PAYPALISHIRING", 4))

}
