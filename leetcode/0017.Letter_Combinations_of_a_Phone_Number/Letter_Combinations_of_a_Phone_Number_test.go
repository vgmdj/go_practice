package _017_Letter_Combinations_of_a_Phone_Number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterCombinations(t *testing.T) {

	ast := assert.New(t)

	ast.EqualValues([]string{"ad", "bd", "cd", "ae", "af", "be", "bf", "ce", "cf"},
		letterCombinations("23"))

	ast.EqualValues([]string{"p", "q", "r", "s"}, letterCombinations("7"))

}
