package leetcode

/*
 * @lc app=leetcode.cn id=844 lang=golang
 *
 * [844] 比较含退格的字符串
 */

// @lc code=start
func backspaceCompare(s string, t string) bool {
	return false
}

func backspace(s string) string {
	if s == "" {
		return ""
	}
	b := []byte(s)
	n := len(b)
	slow := 0
	for fast := 0; fast < n; fast++ {
		if b[slow] != b[fast] {
			b[slow] = b[fast]
			slow++
		}
	}
	return ""
}

// @lc code=end
