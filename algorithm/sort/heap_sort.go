package sort

func HeapSort(values []int) {
	buildHeap(values)
	for i := len(values); i > 1; i-- {
		values[0], values[i-1] = values[i-1], values[0]
		adjustHeap(values[:i-1], 0)
		//fmt.Println("the heap is ", values)
	}
}

func buildHeap(values []int) {
	for i := len(values); i >= 0; i-- { //////一定得从后往前调整，
		adjustHeap(values, i)
	}
}

func adjustHeap(values []int, pos int) { ///////// 调整pos位置的结点
	node := pos
	length := len(values)
	for node < length {
		var child int = 0
		if 2*node+2 < length {
			if values[2*node+1] > values[2*node+2] {
				child = 2*node + 1
			} else {
				child = 2*node + 2
			} ////////选出大子节点
		} else if 2*node+1 < length {
			child = 2*node + 1
		}
		if child > 0 && values[child] > values[node] {
			values[node], values[child] = values[child], values[node]
			node = child
		} else {
			break
		}
	}
}