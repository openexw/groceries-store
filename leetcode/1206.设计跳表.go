/*
 * @lc app=leetcode.cn id=1206 lang=golang
 *
 * [1206] 设计跳表
 */

// @lc code=start
package leetcode

import "math/rand"

type skipNode struct {
	val        int
	next, down *skipNode
}

func newSkipNode(val int, next, down *skipNode) *skipNode {
	return &skipNode{val, next, down}
}

type Skiplist struct {
	head  *skipNode
	len   int
	level int
}

func Constructorx() Skiplist {
	return Skiplist{level: 1, head: newSkipNode(0, nil, nil)}
}

func (this *Skiplist) Search(target int) bool {

}

func (this *Skiplist) Add(num int) {
	level := 1
	if level <= this.level && rand.Int31()&1 == 0 {
		level++
	}

	if level > this.level {
		this.level = level
		this.head = newSkipNode(num, nil, this.head)
	}
	curr := this.head

}

func (this *Skiplist) Erase(num int) bool {

}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
// @lc code=end
