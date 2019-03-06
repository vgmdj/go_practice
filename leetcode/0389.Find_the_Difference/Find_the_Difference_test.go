package Find_the_Difference

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTheDifference(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(byte('e'), findTheDifference("abcd", "ceabd"))
}
