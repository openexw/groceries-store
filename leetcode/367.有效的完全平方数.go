package leetcode

/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	if num == 0 || num == 1 {
		return true
	}
	left, right := 0, num-1
	for right >= left {
		mid := left + (right-left)>>1
		if mid*mid == num {
			return true
		} else if mid*mid > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}

// @lc code=end
