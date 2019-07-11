package Verify_Preorder_Serialization_of_a_Binary_Tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerify(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(false, isValidSerialization("#,#"))
}
