package sort

import (
	"testing"
)

func Test_insertion(t *testing.T) {
	arr := []int64{1, 5, 1, 54, 9}
	data := insertion(arr)
	t.Log(data)
}

func BenchmarkInsertion(b *testing.B) {
	//data := []int64{1, 45, 455656, 5}
	//l := len(data)
	//for i := 1; i < l; i++ {
	//	curr := data[i]
	//	j := i - 1
	//	for ; j > 0; j-- {
	//		if data[j] > curr {
	//			data[j+1] = data[j]
	//		} else {
	//			break
	//		}
	//
	//	}
	//	data[j+1] = curr
	//}
	////t.Log(data)
}
