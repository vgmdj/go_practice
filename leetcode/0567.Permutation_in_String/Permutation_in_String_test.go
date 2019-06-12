package Permutation_in_String

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutation(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(true, checkInclusion("abdacdefg", "defgcdlfgeiadbcdaefggil"))
	ast.Equal(true, checkInclusion("adc", "dcda"))

}
