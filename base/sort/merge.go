package sort

func merge(left, right []int64) []int64 {
	lL := len(left)
	lR := len(right)
	i, j := 0, 0
	tmp := make([]int64, 0)
	for i < lL && j < lR {
		if left[i] < right[j] {
			tmp = append(tmp, left[i])
			i++
		} else {
			tmp = append(tmp, right[j])
			j++
		}
	}
	if i < lL {
		tmp = append(tmp, left[i:]...)
	} else {
		tmp = append(tmp, right[j:]...)
	}
	return tmp
}
func MergeSort(data []int64) []int64 {
	l := len(data)
	if l <= 1 {
		return data
	}

	mid := l / 2
	left := MergeSort(data[:mid])
	right := MergeSort(data[mid:l])
	return merge(left, right)
}

//func mergeS(data []int64, start, end int) []int64 {
//	if start >= end {
//		return data
//	}
//	mid := (start + end) / 2
//	mergeS(data, 0, mid)
//	mergeS(data, mid+1, end)
//
//	merge()
//}
