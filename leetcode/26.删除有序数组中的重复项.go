package leetcode

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	slowIndex := 0
	for fastIndex := 0; fastIndex < length; fastIndex++ {
		if nums[slowIndex] != nums[fastIndex] {
			slowIndex++
			nums[slowIndex] = nums[fastIndex]

		}
	}
	return slowIndex + 1
}

// @lc code=end
