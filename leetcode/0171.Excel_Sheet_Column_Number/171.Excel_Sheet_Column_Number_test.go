package Excel_Sheet_Column_Number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSheetNumber(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(titleToNumber("A"), 1)
	ast.Equal(titleToNumber("AB"), 28)
	ast.Equal(titleToNumber("Z"), 26)
	ast.Equal(titleToNumber("ZY"), 701)
	ast.Equal(titleToNumber("B"), 2)
	ast.Equal(titleToNumber("ZX"), 700)
	ast.Equal(titleToNumber("AAA"), 703)

}
