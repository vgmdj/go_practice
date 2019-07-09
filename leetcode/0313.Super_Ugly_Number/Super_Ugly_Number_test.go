package Super_Ugly_Number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuperUglyNumber(t *testing.T){
	ast := assert.New(t)

	ast.Equal(32,nthSuperUglyNumber(12,[]int{2,7,13,19}))
}
