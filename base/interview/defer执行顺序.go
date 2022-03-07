package interview

import "fmt"

func de() {
	defer func() {
		println("aa")
	}()
	defer func() {
		println("bb")
	}()
	defer func() {
		println("cc")
	}()
	defer func() {
		println("dd")
	}()

	println("ee")
	panic(1)
}

// 1. defer 按照栈结构进行执行。即先进先出
// 2. panic 后 defer 依然会执行
// 3.

func s()  {
	var a string

	//0x1400003df58
	fmt.Printf("%p\n", &a)
}