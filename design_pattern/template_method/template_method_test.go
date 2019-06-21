package template_method

import "testing"

func TestHummerModel_Run(t *testing.T) {
	h1 := NewHummerH1Model()
	h1.Run()

	h2 := NewHummerH2Model()
	h2.Run()
}
