package leetcode

import "testing"

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{"ab#c", "ad#c"}, true},
		{"test2", args{"ab##", "ab##"}, true},
		{"test3", args{"a##c", "a##c"}, true},
		{"test4", args{"a#c", "a#c"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backspace(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{"ab#c"}, "ac"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspace(tt.args.s); got != tt.want {
				t.Errorf("backspace() = %v, want %v", got, tt.want)
			}
		})
	}
}
