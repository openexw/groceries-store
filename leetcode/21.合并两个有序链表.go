/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */
package leetcode

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 方法一：递归
// 时间复杂度：O(M+N)
// 空间复杂度：O(M+N)
//func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//	if list1 == nil {
//		return list2
//	}
//
//	if list2 == nil {
//		return list1
//	}
//
//	if list1.Val < list2.Val {
//		list1.Next = mergeTwoLists(list1.Next, list2)
//		return list1
//	} else {
//		list2.Next = mergeTwoLists(list1, list2.Next)
//		return list2
//	}
//}

// 方法二、暴力解法
// 时间复杂度：O(M+N)
// 空间复杂度：O(1)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 哨兵节点
	head := &ListNode{}
	last := head

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			last.Next = list1
			list1 = list1.Next
		} else {
			last.Next = list2
			list2 = list2.Next
		}
		last = last.Next
	}
	if list1 == nil {
		last.Next = list2
	}
	if list2 == nil {
		last.Next = list1
	}
	return head.Next
}

// @lc code=end
