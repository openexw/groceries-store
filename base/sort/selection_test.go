package sort

import (
	"testing"
)

func Test_selection(t *testing.T) {
	arr := []int64{1, 5, 1, 54, 9}
	data := selection(arr)
	t.Log(data)
}
