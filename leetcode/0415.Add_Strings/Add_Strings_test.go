package Add_Strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddString(t *testing.T) {
	ast := assert.New(t)

	ast.Equal("198",addStrings("99","99"))
	ast.Equal("507",addStrings("6","501"))
	ast.Equal("134", addStrings("121", "13"))
}
