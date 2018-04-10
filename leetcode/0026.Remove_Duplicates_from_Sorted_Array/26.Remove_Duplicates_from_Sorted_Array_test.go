package Remove_Duplicates_from_Sorted_Array

import "testing"

func TestOK(t *testing.T) {
	nums := []int{1, 1, 2, 2, 2, 3, 3}
	t.Log(removeDuplicates(nums))
	t.Log(nums)
}
