package interview

func bubble(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	}

	for i := 0; i < l; i++ {
		flag := false
		for j := 0; j < l-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}

		if !flag {
			break
		}
	}
	return arr
}

func insertion(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	}

	for i := 0; i < l; i++ {
		val := arr[i]

		j := i - 1
		for ; j > 0; j-- {
			if arr[j] > val {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[i+1] = val
	}
	return arr
}

func quicksort(data []int, p, r int) {
	if p > r {
		return
	}
	q := quicksortPartition(data, p, r)
	quicksort(data, p, q-1)
	quicksort(data, q+1, r)
}

func quicksortPartition(data []int, p, r int) int {
	pivot := data[r]
	i := p
	// 将所有小于 pivot 放在左边，并得到最后一个小于 pivot 的位置
	for j := p; j < r; j++ {
		if data[j] < pivot {
			data[j], data[i] = data[i], data[j]
			i++
		}
	}

	// 将最后一个小于 pivot 的位置数据与 pivot 进行交换（使用最后一个值为 pivot）
	data[i], data[r] = data[r], data[i]
	return i
}
