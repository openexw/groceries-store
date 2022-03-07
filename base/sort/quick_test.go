package sort

import (
	"sort"
	"testing"
)

func Test_partition(t *testing.T) {
	arr := []int{6, 11, 3, 9, 8}
	partition(arr, 0, len(arr)-1)
}
func TestQuicksort(t *testing.T) {
	arr := []int{6, 11, 3, 9, 8}

	Quicksort(arr, 0, len(arr)-1)

	if !sort.IntsAreSorted(arr) {
		t.Error("排序数据异常")
	}
}

//func Sort(data []int, p, r int) {
//	if p > r {
//		return
//	}
//
//	q := partitions(data, p, r)
//	Sort(data, p, q-1)
//	Sort(data, q+1, r)
//
//}
//
//func partitions(data []int, p, r int) int {
//	pivot := data[r]
//	i := p
//	l := len(data)
//	for j := p; j < l-1; j++ {
//		if data[j] < pivot {
//			data[j], data[i] = data[i], data[j]
//			i++
//		}
//	}
//
//	data[i], data[r] = data[r], data[i]
//	return i
//}
