package sort

func Bubble(data []int64) []int64 {
	l := len(data)
	// 如果数据长度 <= 1 直接返回
	if l <= 1 {
		return data
	}
	for i := 0; i < l; i++ { // 需要冒泡多少趟
		// 提前退出冒泡循环的标志位
		flag := false // 当某次冒泡操作已经没有数据交换时，说明已经达到完全有序，不用再继续执行后续的冒泡操作
		for j := 0; j < l-1-i; j++ { // 对当前排序区做自上而下的扫描
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				flag = true // 数据发生了交换
			}
		}
		// 没有数据交换，提前退出
		if !flag {
			break
		}
	}
	return data
}
