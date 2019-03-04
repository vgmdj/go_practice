package Divide_Two_Integers

import (
	"math/rand"
	"testing"
)

func TestDivide(t *testing.T) {
	for i := 0; i < 100; i++ {
		a := rand.Intn(10000)
		b := rand.Intn(1000)
		c := divide(a, b)

		if c != a/b {
			t.Error(a, b, a/b, c)
			return
		}

	}

}
