package sort

func Quicksort(data []int, p, r int) {
	if p > r {
		return
	}

	q := partition(data, p, r)

	Quicksort(data, p, q-1)
	Quicksort(data, q+1, r)
}

func partition(data []int, p, r int) int {
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
