package sorting

func HeapSort(list []int) {
	for start := (len(list) - 2) / 2; start >= 0; start-- {
		siftDown(list, start, len(list)-1)
	}

	for end := len(list) - 1; end > 0; end-- {
		swap(list, 0, end)
		siftDown(list, 0, end-1)
	}
}

func siftDown(list []int, start, end int) {
	for root := start; root*2+1 <= end; {
		child := root*2 + 1
		if child+1 <= end && list[child] < list[child+1] {
			child++
		}

		if list[root] >= list[child] {
			return
		}
		swap(list, root, child)
		root = child
	}
}

func swap(list []int, i, j int) {
	list[i], list[j] = list[j], list[i]
}
