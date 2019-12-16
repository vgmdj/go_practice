package String_to_Integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtoi(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(myAtoi("   010"), 10)
	ast.Equal(myAtoi("001230"), 1230)
	ast.Equal(myAtoi("+1230"), 1230)
	ast.Equal(myAtoi("-1230"), -1230)
	ast.Equal(myAtoi("616303511230"), 2147483647)
	ast.Equal(myAtoi("  -0012a42"), -12)
	ast.Equal(myAtoi("-2147483648"), -2147483648)
	ast.Equal(myAtoi("9223372036854775809"), 2147483647)
	ast.Equal(myAtoi("-6147483648"), -2147483648)

}
