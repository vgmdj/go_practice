package sort

/*
快速排序
假设初始数组为
5 4 3 2 1 8 9 0
h n           t
m

如果和选定中间值比较,
大于m, 与t上的值交换, 并使t前移
小于等于m, 与上一位置上的值交换, 即与h上的值交换, 并将h与n后移
(所以for循环条件可写成 head < tail 或者 next <= tail)
h实际代表的是选定中间值的位置

0 1 2 3 4 5 6 7 (p)
4 3 2 1 0 5 9 8	(v)
          h n
          t
所以, 5的位置已经排好, 还需要排 0-4 和 6-7的位置
quickSort(values[:n-1])
quickSort(values[n:])

*/
func quickSort(values []int) {
	if len(values) <= 1 {
		return
	}
	midValue, next := values[0], 1
	head, tail := 0, len(values)-1
	for head < tail {
		if values[next] > midValue {
			values[next], values[tail] = values[tail], values[next]
			tail--
		} else {
			values[next], values[head] = values[head], values[next]
			head++
			next++
		}
	}

	quickSort(values[:next-1])
	quickSort(values[next:])

}
