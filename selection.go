package sorting

func SelectionSort(list []int) {
	for i := range list {
		k := i
		for j := i + 1; j < len(list); j++ {
			if list[j] < list[k] {
				k = j
			}
			list[i], list[k] = list[k], list[i]
		}
	}
}
