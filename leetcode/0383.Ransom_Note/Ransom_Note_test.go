package _383_Ransom_Note

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRansomNote(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(false, canConstruct("az", "bcz"))
}
