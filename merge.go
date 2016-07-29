package sorting

func MergeSort(list []int) {
	s := make([]int, len(list)/2+1)
	mergeSort(list, s)
}

func mergeSort(list []int, s []int) {
	if len(list) < 2 {
		return
	}

	mid := len(list) / 2
	mergeSort(list[:mid], s)
	mergeSort(list[mid:], s)

	if list[mid-1] <= list[mid] {
		return
	}

	copy(s, list[:mid])
	l, r := 0, mid
	for i := 0; ; i++ {
		if s[l] <= list[r] {
			list[i] = s[l]
			l++
			if l == mid {
				break
			}
		} else {
			list[i] = list[r]
			r++
			if r == len(list) {
				copy(list[i+1:], s[l:mid])
				break
			}
		}
	}
}
