package sort

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int64{9, 10, 0, 18, 1, -9, 4, 45}
	sort := MergeSort(arr)
	t.Log(sort)
}
