package Range_Sum_Query_Immutable

type NumArray struct {
	Sums []int
}

func Constructor(nums []int) NumArray {
	na := NumArray{}
	na.Sums = make([]int, len(nums)+1)

	for k,v := range nums {
		na.Sums[k+1] = na.Sums[k] + v
	}

	return na

}

func (this *NumArray) SumRange(i int, j int) int {
	return this.Sums[j+1] - this.Sums[i]
}
