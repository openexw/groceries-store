package leetcode

import "testing"

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test1", args{[]int{0, 1, 0, 3, 12}}},
		{"test2", args{[]int{0, 1, 1}}},
		{"test2", args{[]int{1, 1, 13}}},
		{"test2", args{[]int{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
		})
	}
}
