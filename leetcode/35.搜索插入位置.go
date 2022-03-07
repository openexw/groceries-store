package leetcode

/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * 分为 4 中情况讨论
 * 1. target 在 nums 所有元素之前
 * 2. target 等于 nums 其中的一个元素
 * 3. target 在 nums 插入其中一个元素
 * 4. target 在 nums 所有元素之后
 * [35] 搜索插入位置
 */
// @lc code=start
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for right >= left {
		mid := left + (right-left)>>1
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return right + 1
}

// @lc code=end
