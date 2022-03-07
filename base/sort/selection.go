package sort

func selection(data []int64) []int64 {
	l := len(data)
	if l <= 1 {
		return data
	}

	for i := 0; i < l-1; i++ {
		min := i

		for j := l - 1; j > i; j-- {
			if data[j] < data[min] {
				min = j
			}
		}

		// 有序则不需要交换
		if i != min {
			data[i], data[min] = data[min], data[i]
		}
	}

	return data
}
