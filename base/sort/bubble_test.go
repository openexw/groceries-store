package sort

import (
	"testing"
)

//curl -XPOST "http://dev-pschool-crm.myyixue.com/api/dingding/student/queryDropDownClassList" -H "content-type:application/json" -H "authorization:626040105194323968" -d '{"campusIds":[2394]}'
func TestBubble(t *testing.T) {
	arr := []int64{
		3, 4, 10, 0, 12,
	}

	data := Bubble(arr)
	t.Log(data)
}

type int6 []int64
type data struct {
	d []int6
}

func BenchmarkBubble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int64{
			3, 4, 10, 0, 12,
		}
		Bubble(arr)
	}
}
