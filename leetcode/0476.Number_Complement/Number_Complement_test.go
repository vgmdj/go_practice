package Number_Complement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberComplement(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(2, findComplement(5))
	ast.Equal(1, findComplement(0))
	ast.Equal(0, findComplement(1))

}
