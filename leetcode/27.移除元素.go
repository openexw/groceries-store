package leetcode

/*
 * @lc app=leetcode.cn id=27 lang=golang
 * 数组的元素在内存地址中是连续的，不能单独删除数组中的某个元素，只能覆盖。
 * 解决方案：
 * - 暴力破解
 * - 快慢指针
 * [27] 移除元素
 */

// @lc code=start
// removeElement 暴力破解法，时间复杂度 O(n^2)，空间复杂度 O(1)
// func removeElement(nums []int, val int) int {
// 	size := len(nums)
// 	for i := 0; i < size; i++ {
// 		// 如果元素相等，需要将后面的元素往前移动
// 		if nums[i] == val {
// 			// i 后面的元素往前移动
// 			for j := i + 1; j < size; j++ {
// 				nums[j-1] = nums[j]
// 			}
// 			// 数据往前移动，i 也往前移动一位
// 			i--
// 			// 数据往前移动，数据出现覆盖，总长度减一
// 			size--
// 		}
// 	}
// 	return size
// }

// removeElement 快慢指针
func removeElement(nums []int, val int) int {
	slowIndex := 0
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != val {
			nums[slowIndex], nums[fastIndex] = nums[fastIndex], nums[slowIndex]
			slowIndex++
		}
	}
	return slowIndex
}

// @lc code=end
