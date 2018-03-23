package Roman_to_Integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoman(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(romanToInt("IV"), 4)
	ast.Equal(romanToInt("XXXIX"), 39)
	ast.Equal(romanToInt("MDCCCLXXXVIII"), 1888)
	ast.Equal(romanToInt("MCMLXXVI"), 1976)
	ast.Equal(romanToInt("MMMCMXCIX"), 3999)

}

func TestRomanII(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(romanToIntII("IV"), 4)
	ast.Equal(romanToIntII("XXXIX"), 39)
	ast.Equal(romanToIntII("MDCCCLXXXVIII"), 1888)
	ast.Equal(romanToIntII("MCMLXXVI"), 1976)
	ast.Equal(romanToIntII("MMMCMXCIX"), 3999)

}

func TestRomanIII(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(romanToIntIII("IV"), 4)
	ast.Equal(romanToIntIII("XXXIX"), 39)
	ast.Equal(romanToIntIII("MDCCCLXXXVIII"), 1888)
	ast.Equal(romanToIntIII("MCMLXXVI"), 1976)
	ast.Equal(romanToIntIII("MMMCMXCIX"), 3999)

}
