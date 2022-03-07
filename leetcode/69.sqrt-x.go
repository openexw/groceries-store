package leetcode

/*
 * @lc app=leetcode.cn id=69 lang=golang
 * 1. 开平方根后是整数的
 * 2. 开平方根后是小数的，结果只保留 整数部分 ，小数部分将被 舍去
 * 利用二分法一个一个的去变量，时间福再度 O(log n)
 * [69] Sqrt(x)
 */
// @lc code=start
func mySqrt(x int) int {
	// 处理特殊情况
	if x == 0 || x == 1 {
		return x
	}
	left, right := 0, x-1
	for right >= left {
		mid := left + (right-left)>>1
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

// @lc code=end
