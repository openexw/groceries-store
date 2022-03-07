package sort

func insertion(data []int64) []int64 {
	l := len(data)
	if l <= 1 {
		return data
	}

	for i := 1; i < l; i++ {
		val := data[i]

		j := i - 1
		// 已排序区间
		for ; j > 0; j-- {
			if data[j] > val {
				// 移动元素
				data[j+1] = data[j]
			} else {
				break
			}
		}
		// 经过在已排序区间元素移动，得到了当前元素 val 所需奥插入的位置 j，并赋值
		data[j+1] = val
	}
	return data
}
