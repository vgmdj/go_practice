package Happy_Number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHappy(t *testing.T) {
	ast := assert.New(t)
	ast.Equal(isHappy(19), true)
	ast.Equal(isHappy(1), true)
	ast.Equal(isHappy(2), false)
	ast.Equal(isHappy(0), false)
}
