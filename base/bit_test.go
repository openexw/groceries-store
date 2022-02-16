package base

import "testing"

const (
	a  = 10000998811000
	a2 = 12
)

func BenchmarkMod1(b *testing.B) {
	//a := 17
	//a2 := 6
	for i := 0; i < b.N; i++ {
		mod1(a, a2)
	}
	b.Logf("%v", uintptr(1))
}

func BenchmarkMod2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mod2(a, a2)
	}
}
