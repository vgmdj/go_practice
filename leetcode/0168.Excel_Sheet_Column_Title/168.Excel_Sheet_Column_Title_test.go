package Excel_Sheet_Column_Title

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSheetTitle(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(convertToTitle(1), "A")
	ast.Equal(convertToTitle(28), "AB")
	ast.Equal(convertToTitle(26), "Z")
	ast.Equal(convertToTitle(701), "ZY")
	ast.Equal(convertToTitle(2), "B")
	ast.Equal(convertToTitle(700), "ZX")
	ast.Equal(convertToTitle(703), "AAA")

}
