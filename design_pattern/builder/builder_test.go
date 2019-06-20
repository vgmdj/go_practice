package builder

import "testing"

func TestBuilder(t *testing.T) {
	tb := new(ThinBuilder)
	d := new(Director)
	d.SetBuilder(tb)
	d.Build()

}
