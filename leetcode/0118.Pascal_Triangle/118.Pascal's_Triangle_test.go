package Pascal_s_Triangle

import "testing"

func TestGeneral(t *testing.T) {
	t.Log(generate(0))
	t.Log(generate(1))
	t.Log(generate(2))
	t.Log(generate(3))
	t.Log(generate(4))
	t.Log(generate(5))
}
