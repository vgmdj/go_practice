package singleton

import "testing"

func TestSingleton(t *testing.T) {
	a := GetInstance()
	b := GetInstance()

	if a != b {
		t.Errorf("different")
	}

}
