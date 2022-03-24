/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */
package leetcode

import (
	"container/list"
)

// @lc code=start
type LRUCache struct {
	list.List
	items map[int]*list.Element
	cap   int
}

func Constructor1(capacity int) LRUCache {
	return LRUCache{
		items: make(map[int]*list.Element, capacity),
		cap:   capacity,
	}
}

type item struct {
	key   int
	value interface{}
}

func newItem(key int, value interface{}) item {
	return item{key: key, value: value}
}

// func (this *LRUCache) Get(key int) int {
// 	if e, ok := this.items[key]; ok {
// 		// move
// 		this.MoveToBack(e)
// 		return e.Value.(*item).value.(int)
// 	}
// 	return -1
// }

// func (this *LRUCache) Put(key int, value int) {
// 	if e, ok := this.items[key]; ok {
// 		// 存在则替换
// 		i := e.Value.(*item)
// 		i.value = value
// 		e.Value = i
// 		this.MoveToBack(e)
// 	} else {
// 		if this.Len() >= this.cap {
// 			// 删除 head 元素
// 			head := this.Front()
// 			this.Remove(head)
// 			// 删除 map
// 			i := head.Value.(*item)
// 			delete(this.items, i.key)
// 		}
// 		v := newItem(key, value)
// 		newW := this.PushBack(v)
// 		this.items[key] = newW
// 	}
// }

func (c *LRUCache) Get(key int) int {
	if e, ok := c.items[key]; ok {
		c.MoveToBack(e)
		return e.Value.(item).value.(int)
	}
	return -1
}

func (c *LRUCache) Put(key, value int) {
	if e, ok := c.items[key]; ok {
		i := e.Value.(item)
		i.value = value
		e.Value = i
		c.MoveToBack(e)
	} else {
		if c.Len() >= c.cap {
			head := c.Front()
			c.Remove(head)

			i := head.Value.(item)
			delete(c.items, i.key)
		}
		v := newItem(key, value)
		newItem := c.PushBack(v)
		c.items[key] = newItem
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
