package Sum_of_Two_Integers

import (
	"math/rand"
	"testing"
)

func TestSum(t *testing.T) {
	for i := 0; i < 100; i++ {
		a := rand.Int()
		b := rand.Int()
		if a+b != getSum(a, b) {
			t.Error(a, b, getSum(a, b))
			return
		}
	}

}
