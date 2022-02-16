package base

func mod1(a, b int) int {
	return a % b
}

// mod2 取余
func mod2(a, b int) int {
	return a & (b - 1)
}

// pow2 b^2
func pow2(b int) int {
	return 1 << b
	//return uintptr(1)<<b - 1
}
