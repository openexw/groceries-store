package leetcode

/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	isEnd    bool
	children map[rune]*Trie
}

func Constructor() Trie {
	return Trie{isEnd: false, children: make(map[rune]*Trie)}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		if child, ok := node.children[ch]; ok {
			node = child
		} else {
			child := &Trie{false, make(map[rune]*Trie)}
			node.children[ch] = child
			node = child
		}
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for _, ch := range word {
		if child, ok := node.children[ch]; ok {
			node = child
			continue
		}
		return false
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, ch := range prefix {
		if child, ok := node.children[ch]; ok {
			node = child
			continue
		}
		return false
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
