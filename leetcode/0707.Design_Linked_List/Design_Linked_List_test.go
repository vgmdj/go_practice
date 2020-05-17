package Design_Linked_List

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	obj := Constructor()

	obj.AddAtHead(1)
	obj.DeleteAtIndex(0)
	obj.AddAtTail(3)
	obj.DeleteAtIndex(0)
	obj.AddAtIndex(-1, 1)
	obj.AddAtIndex(1, 3)
	obj.AddAtIndex(1, 2)

	val := obj.Get(1)
	if val != 2 {
		t.Errorf("want to get 2 ,but get %v", val)
		return
	}

	obj.DeleteAtIndex(1)
	val = obj.Get(1)
	if val != 3 {
		t.Errorf("want to get 3 ,but get %v", val)
		return
	}

}
