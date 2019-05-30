package Sliding_Window_Maximum

func maxSlidingWindow(nums []int, k int) []int {
	windows := []int{}
	res := []int{}
	// 策略修改了, window append 进去的是index 序列号, 而不是具体的值
	// window[0] 负责维护窗口最大那个
	for i, _ := range nums {
		if i >= k && windows[0] <= i-k {
			// 不是前两个 && 不是最后两个, 就推出第一个值,
			windows = windows[1:]
		}
		for len(windows) > 0 && nums[windows[len(windows)-1]] < nums[i] {
			// 窗口内 当前要被从窗口后面推进去的那个数字, 它的前一个数字比他小, 那他前面那个数字推出来.
			windows = windows[:len(windows)-1]
		}
		// 推进去当前 序列 i 的数字进入窗口
		windows = append(windows, i)
		if i > k-2 {
			// 从第三个开始
			res = append(res, nums[windows[0]])
		}
	}
	return res
}