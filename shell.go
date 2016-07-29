package sorting

func ShellSort(list []int) {
	for inc := len(list) / 2; inc > 0; inc = (inc + 1) * 5 / 11 {
		for i := inc; i < len(list); i++ {
			j, temp := i, list[i]
			for ; j >= inc && list[j-inc] > temp; j -= inc {
				list[j] = list[j-inc]
			}
			list[j] = temp
		}
	}
}
