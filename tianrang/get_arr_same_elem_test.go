package tianrang

import (
	"reflect"
	"testing"
)

func Test_arrSameElem(t *testing.T) {
	type args struct {
		a1 []int
		a2 []int
	}
	tests := []struct {
		name    string
		args    args
		wantRst []int
	}{
		{
			"test_1",
			args{
				a1: []int{1, 6, 2, 1},
				a2: []int{2, 2},
			},
			[]int{2, 2},
		},
		{
			"test_2",
			args{
				a1: []int{7, 3, 5},
				a2: []int{3, 7, 3, 8, 7},
			},
			[]int{3, 7, 3, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRst := arrSameElem(tt.args.a1, tt.args.a2); !reflect.DeepEqual(gotRst, tt.wantRst) {
				t.Errorf("arrSameElem() = %v, want %v", gotRst, tt.wantRst)
			}
		})
	}
}
