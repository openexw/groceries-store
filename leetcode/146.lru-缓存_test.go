package leetcode

import (
	"fmt"
	"testing"
)

func TestConstructor1(t *testing.T) {
	l := Constructor1(2)
	l.Put(1, 1)
	l.Put(2, 2)
	l.Get(1)
	l.Put(3, 3)
	l.Get(2) // -1
	l.Put(4, 4)
	l.Get(1)
	l.Get(3)
	l.Get(4)
	fmt.Printf("%v\n", l.items)
}

func TestLRUCache_Put(t *testing.T) {
	//l := Constructor1(2)
	//l.Put(1, 1)   // [1]
	//l.Put(2, 2)   // [1,2]
	//r := l.Get(1) // 1 [2,1]
	//want := 1
	//if r != want {
	//	t.Errorf("go result = %d,but got is %d\n", r, want)
	//}
	//
	//l.Put(3, 3)  // [2x,1,3] => [1,3]
	//r = l.Get(2) // -1
	//want = -1
	//if r != want {
	//	t.Errorf("go result = %d,but got is %d\n", r, want)
	//}
	//
	//l.Put(4, 4)  // [1x,3, 4] => [3,4]
	//r = l.Get(1) // -1
	//want = -1
	//if r != want {
	//	t.Errorf("go result = %d,but got is %d\n", r, want)
	//}
	//
	//r = l.Get(3)
	//want = 3
	//if r != want {
	//	t.Errorf("go result = %d,but got is %d\n", r, want)
	//}
	//r = l.Get(4)
	//want = 4
	//if r != want {
	//	t.Errorf("go result = %d,but got is %d\n", r, want)
	//}

	//fmt.Printf("%v\n", l.items)

	l := Constructor1(1)
	l.Put(2, 1)
	l.Get(2)
}
