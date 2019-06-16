package Third_Maximum_Number

// O(3n)
func thirdMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	tm := ThirdNumber{TopThree: make([]int, 0)}

	for _, v := range nums {
		tm.Push(v)
	}

	return tm.Result()

}

type ThirdNumber struct {
	TopThree []int
}

func (t *ThirdNumber) Push(x int) {
	for _, v := range t.TopThree {
		if x == v {
			return
		}
	}

	if len(t.TopThree) < 3 {
		t.TopThree = append(t.TopThree, x)
		t.Sort()
		return
	}

	if t.TopThree[2] > x {
		return
	}

	t.TopThree[2] = x
	t.Sort()

}

func (t *ThirdNumber) Sort() {
	for i := 0; i < len(t.TopThree); i++ {
		max := 0

		for j := i + 1; j < len(t.TopThree); j++ {
			if t.TopThree[max] < t.TopThree[j] {
				max = j
			}
		}

		if i != max {
			t.TopThree[i], t.TopThree[max] = t.TopThree[max], t.TopThree[i]
		}

	}

}

func (t *ThirdNumber) Result() int {
	if len(t.TopThree) < 3 {
		return t.TopThree[0]
	}

	return t.TopThree[2]

}
