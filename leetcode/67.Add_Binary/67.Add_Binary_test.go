package Add_Binary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(addBinary("1", "1"), "10")
	ast.Equal(addBinary("11", "1"), "100")
	ast.Equal(addBinary("10", "10"), "100")
}
