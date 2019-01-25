package Integer_to_Roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoman(t *testing.T) {
	ast := assert.New(t)

	ast.Equal("III", intToRoman(3))
	ast.Equal("IV", intToRoman(4))
	ast.Equal("V", intToRoman(5))
	ast.Equal("VI", intToRoman(6))
	ast.Equal("IX", intToRoman(9))
	ast.Equal("MCMXCIV", intToRoman(1994))

}
