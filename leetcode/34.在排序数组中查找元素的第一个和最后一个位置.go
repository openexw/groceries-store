package leetcode
/*
 * @lc app=leetcode.cn id=34 lang=golang
 * 1.查找第一个值等于给定值的元素
 * 2.查找最后一个值等于给定值的元素
 * 3.查找第一个大于等于给定值的元素
 * 4.查找最后一个小于等于给定值的元素
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */
// @lc code=start
func searchRange(nums []int, target int) []int {
	left := binarySearchStart(nums, target)
	right := binarySearchEnd(nums, target)
	return []int{left, right}
	//if left == -1 || right == -1 {
	//	return []int{-1, -1}
	//}
	//
	//if right-left > 1 {
	//	return []int{left + 1, right - 1}
	//}
	//return []int{-1, -1}
}

//
//func binarySearchStart(nums []int, target int) int {
//	left, right := 0, len(nums)-1
//	pos := -1
//	for left <= right {
//		mid := left + (right-left)/2
//		if nums[mid] >= target {
//			right = mid - 1
//			pos = right
//		} else {
//			left = mid + 1
//
//		}
//	}
//	return pos
//}
//func binarySearchEnd(nums []int, target int) int {
//	left, right := 0, len(nums)-1
//	pos := -1
//	for left <= right {
//		mid := left + (right-left)/2
//		if nums[mid] > target {
//			right = mid - 1
//		} else {
//			left = mid + 1
//			pos = left
//		}
//	}
//	return pos
//}

func binarySearchStart(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for right >= left {
		mid := left + (right-left)>>1
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			if mid == 0 || nums[mid-1] != target {
				return mid
			}
			right = mid - 1
		}
	}
	return -1
}
func binarySearchEnd(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for right >= left {
		mid := left + (right-left)>>1
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			if (mid == len(nums)-1) || nums[mid+1] != target {
				return mid
			}
			left = mid + 1
		}
	}
	return -1
}

// @lc code=end
