package strategy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiscount(t *testing.T) {
	ast := assert.New(t)

	discount := NewCashContext("normal")
	ast.Equal(discount.Accept(300), float64(300))

	discount = NewCashContext("0.8rebate")
	ast.Equal(discount.Accept(300), float64(240))

	discount = NewCashContext("300return100")
	ast.Equal(discount.Accept(300), float64(200))

}
