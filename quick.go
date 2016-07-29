package sorting

func QuickSort(list []int) {
	sort(list, 0, len(list)-1)
}

func sort(list []int, low, high int) {
	i := low
	j := high
	pivot := list[low+(high-low)/2]

	for i <= j {
		for list[i] < pivot {
			i++
		}

		for list[j] > pivot {
			j--
		}

		if i <= j {
			swap(list, i, j)
			i++
			j--
		}
	}

	if low < j {
		sort(list, low, j)
	}

	if i < high {
		sort(list, i, high)
	}
}
