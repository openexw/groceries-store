package leetcode

import (
	"testing"
)

func TestConstructor2(t *testing.T) {
	//l := Constructor2()
	//l.AddAtHead(7)
	//l.AddAtHead(2)
	//l.AddAtHead(1) // [1, 2, 7]
	//l.print()
	//l.AddAtIndex(3, 0) // [1, 2, 7, 0]
	//l.print()
	//l.DeleteAtIndex(2)

	// [24,30]
	l := Constructor2()
	l.AddAtHead(1)
	l.DeleteAtIndex(0)
	l.print()
	//l.AddAtHead(24)
	//l.Get(1) // -1
	//l.AddAtTail(18)
	//l.DeleteAtIndex(1)
	//l.Get(1) // -1
	//l.AddAtTail(30)
	//println(l.Get(2)) // -1
	//l.print() // [24 30]
	//l.DeleteAtIndex(1)
	//l.AddAtHead(3)
	//l.AddAtHead(3)
	//l.AddAtHead(33)
	//l.AddAtHead(97)
	//l.AddAtTail(43)
	//l.AddAtTail(12)
	//l.AddAtTail(10)
	//l.DeleteAtIndex(1)
	//l.AddAtIndex(1,56)
	//l.AddAtHead(30)
	//l.AddAtIndex(8,83)
	//l.AddAtTail(57)
	//l.AddAtHead(74)
	//l.Get(5)
	//l.AddAtHead(98)
	//l.AddAtTail(72)
	//l.AddAtHead(34)
	//l.AddAtHead(61)
	//l.Get(6)
	//l.AddAtHead(70)
	//l.AddAtHead(24)
	//l.AddAtTail(91)
	//l.AddAtHead(99)
	//l.AddAtTail(13)
	//l.AddAtTail(10)
	//l.DeleteAtIndex(17)
	//l.AddAtTail(84)
	//l.DeleteAtIndex(16)
	//l.AddAtHead(73)
	//l.AddAtTail(88)
	//l.AddAtIndex(4,19)
	//l.AddAtHead(59)
	//l.AddAtTail(41)
	//l.AddAtHead(57)
	//l.DeleteAtIndex(10)
	//l.DeleteAtIndex(18)
	//l.AddAtHead(2)
	//l.AddAtTail(12)
	//l.AddAtTail(25)
	//l.AddAtHead(1)
	//l.AddAtHead(77)
	//l.Get(1)
	//l.DeleteAtIndex(7)
	//l.AddAtHead(34)
	//l.AddAtTail(87)
	//l.DeleteAtIndex(13)
	//l.AddAtTail(4)
	//l.AddAtTail(12)
	//l.AddAtTail(11)
	//l.AddAtIndex(10,92)
	//l.AddAtIndex(21,55)
	//l.Get(11)
	//l.AddAtTail(38)
	//l.AddAtTail(31)
	//l.AddAtHead(45)
	//l.AddAtTail(4)
	//l.AddAtTail(21)
	//l.AddAtHead(38)
	//l.Get(4)
	//l.AddAtTail(88)
	//l.Get(12)
	//l.DeleteAtIndex(22)
	//l.AddAtHead(40)
	//l.AddAtHead(22)
	//l.AddAtHead(23)
	//l.AddAtIndex(13,96)
	//l.AddAtIndex(24,50)
	//l.DeleteAtIndex(8)
	//l.Get(14)
	//l.AddAtTail(25)
	//l.AddAtTail(53)
	//l.AddAtHead(42)
	//l.Get(6)
	//l.AddAtHead(58)
	//l.AddAtTail(55)
	//l.AddAtIndex(18,72)
	//l.DeleteAtIndex(13)
	//l.AddAtHead(30)
	//l.AddAtHead(97)
	//l.AddAtTail(59)
	//l.Get(47)
	//l.DeleteAtIndex(24)
	//l.AddAtHead(37)
	//l.AddAtTail(26)
	//l.AddAtTail(31)
	//l.AddAtHead(93)
	//l.AddAtHead(66)
	//l.DeleteAtIndex(11)
	//l.Get(43)
	//l.AddAtHead(70)
	//l.AddAtTail(36)
	//l.AddAtHead(31)
	//l.AddAtTail(28)
}
//
//func TestMyLinkedList_AddAtHead(t *testing.T) {
//	l := Constructor2()
//	l.AddAtHead(1)
//	l.AddAtHead(2)
//	l.AddAtHead(3)
//
//	want := []int{3, 2, 1}
//	r := l.iter()
//	if !reflect.DeepEqual(r, want) || l.len != len(want) {
//		t.Errorf("go rst = %v,but want = %v\n", r, want)
//	}
//}
//
//func TestMyLinkedList_AddAtTail(t *testing.T) {
//	l := Constructor2()
//	l.AddAtTail(1)
//	l.AddAtTail(2)
//	l.AddAtTail(3)
//
//	want := []int{1, 2, 3}
//	r := l.iter()
//	if !reflect.DeepEqual(r, want) || l.len != len(want) {
//		t.Errorf("go rst = %v,but want = %v\n", r, want)
//	}
//}
//
//func TestMyLinkedList_AddAtIndex(t *testing.T) {
//	l := Constructor2()
//	l.AddAtTail(1)
//	l.AddAtTail(2)
//	l.AddAtTail(3)
//	// [1,2,3]
//	l.AddAtIndex(0, 12) // [12, 1, 2, 3]
//	r := l.Get(1)
//	want := 1
//	if r != want {
//		t.Errorf("go rst = %v,but want = %v\n", r, want)
//	}
//
//	l.AddAtIndex(4, 10) // [12, 1, 2, 3, 10]
//	r = l.Get(4)
//	//l.print()
//	want = 10
//	if r != want {
//		t.Errorf("go rst = %v,but want = %v\n", r, want)
//	}
//
//	l.AddAtIndex(2, 4) // [12, 1, 4, 2, 3, 10]
//	rs := l.iter()
//	w := []int{12, 1, 4, 2, 3, 10}
//	if !reflect.DeepEqual(rs, w) {
//		t.Errorf("go rst = %v,but want = %v\n", rs, w)
//	}
//}
//
//func TestMyLinkedList_DeleteAtIndex(t *testing.T) {
//	l := Constructor2()
//	l.AddAtTail(1)
//	l.AddAtTail(2)
//	l.AddAtTail(3)
//	// [1,2,3]
//	l.AddAtIndex(0, 12) // [12, 1, 2, 3]
//	l.AddAtIndex(0, 11) // [11, 12, 1, 2, 3]
//	l.AddAtIndex(0, 10) // [10, 11, 12, 1, 2, 3]
//
//	l.DeleteAtIndex(0) // [11, 12, 1, 2, 3]
//	want := []int{11, 12, 1, 2, 3}
//	r := l.iter()
//	l.print()
//	if !reflect.DeepEqual(r, want) || l.len != len(want) {
//		t.Errorf("go rst = %v,but want = %v\n", r, want)
//	}
//
//	l.DeleteAtIndex(4) // [11, 12, 1, 2]
//	want = []int{11, 12, 1, 2}
//	r = l.iter()
//	l.print()
//	if !reflect.DeepEqual(r, want) || l.len != len(want) {
//		t.Errorf("go rst = %v,but want = %v\n", r, want)
//	}
//	l.DeleteAtIndex(1) // [11, 1, 2]
//	want = []int{11, 1, 2}
//	r = l.iter()
//	l.print()
//	if !reflect.DeepEqual(r, want) || l.len != len(want) {
//		t.Errorf("go rst = %v,but want = %v\n", r, want)
//	}
//}
