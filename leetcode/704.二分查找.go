/*
 * @lc app=leetcode id=704 lang=golang
 *
 * [704] Binary Search
 */
package leetcode

// @lc code=start
// search plan 1
// func search(nums []int, target int) int {
// 	for i, num := range nums {
// 		if num == target {
// 			return i
// 		}
// 	}
// 	return -1
// }

// search use [left, right] 的思路，左闭右闭，即 target 包含在 [left, right] 内
// func search(nums []int, target int) int {
// 	left := 0
// 	right := len(nums) - 1
// 	for left <= right {
// 		middle := left + (right-left)/2
// 		if nums[middle] > target {
// 			right = middle - 1
// 		} else if nums[middle] < target {
// 			left = middle + 1
// 		} else {
// 			return middle
// 		}
// 	}
// 	return -1
// }

// 使用 [left, right) 的思路，左闭右开，即 target 不包含在 [left, right) 中
func search(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		// 防止溢出 等同于(left + right)/2
		middle := left + (right-left)/2
		if nums[middle] > target {
			right = middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			return middle
		}
	}
	return -1
}


// @lc code=end
