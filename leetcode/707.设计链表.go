/*
 * @lc app=leetcode.cn id=707 lang=golang
 *
 * [707] 设计链表
 */
package leetcode

import "fmt"

// @lc code=start

type LinkedNode struct {
	value int
	prev  *LinkedNode
	next  *LinkedNode
}
type MyLinkedList struct {
	len        int
	head, tail *LinkedNode
}

func Constructor2() MyLinkedList {
	return MyLinkedList{head: nil, tail: nil}
}

func (this *MyLinkedList) get(index int) *LinkedNode {
	node := this.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index > this.len {
		return -1
	}

	node := this.get(index)
	if node == nil {
		return -1
	}
	return node.value
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &LinkedNode{
		value: val,
		prev:  nil,
		next:  this.head,
	}
	if this.head == nil {
		this.head = newNode
		this.tail = newNode
	} else {
		this.head.prev = newNode
		//this.tail = this.head
		this.head = newNode
	}
	this.len++
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &LinkedNode{
		value: val,
		prev:  this.tail,
		next:  nil,
	}
	if this.tail != nil {
		this.tail.next = newNode
		this.tail = newNode
	} else {
		this.tail = newNode
		this.head = newNode
	}
	this.len++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}

	if this.len == index {
		this.AddAtTail(val)
		return
	}
	if index > this.len {
		return
	}

	node := this.get(index)
	newNode := &LinkedNode{
		value: val,
		prev:  node.prev,
		next:  node,
	}

	//原节点
	node.prev.next = newNode
	node.prev = newNode
	this.len++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.len || this.len == 0 {
		return
	}
	//if index == 0 {
	//	this.head = this.head.next
	//	//this.head.prev = nil
	//} else if index == this.len-1 {
	//	this.tail = this.tail.prev
	//	//this.tail.next = nil
	//} else {
	//	node := this.get(index)
	//
	//	node.next.prev = node.prev
	//	node.prev.next = node.next
	//}
	node := this.get(index)
	if node.next == nil && node.prev == nil {
		this.head = nil
		this.tail = nil
	} else if node.next == nil {
		this.tail = this.tail.prev
		this.tail.next = nil
	} else if node.prev == nil {
		// 头结点
		this.head = this.head.next
		this.head.prev = nil
	} else {
		node.next.prev = node.prev
		node.prev.next = node.next
	}
	this.len--
}

func (this *MyLinkedList) iter() []int {
	var a []int
	for i := 0; i < this.len; i++ {
		a = append(a, this.Get(i))
	}
	return a
}
func (this *MyLinkedList) print() {
	fmt.Printf("%v\n", this.iter())
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
// @lc code=end
