package unan

import "strings"

type node struct {
	path string
	//nType     int8
	//isEnd     bool
	children  []*node
	wildChild bool // 是否是精准匹配，包含 `：`、`*` 等为 true
	fullPath  string
}

func newNode() *node {
	return &node{children: make([]*node, 0)}
}

func parsePath(path string) []string {
	vs := strings.Split(path, "/")
	var parts []string
	for _, part := range vs {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (n *node) insertChild(path string) {
	parts := parsePath(path)
	parent := n
	for _, part := range parts {
		child := n.matchChild(part)
		if child != nil {
			parent = child
			continue
		}
		child = &node{path: part, wildChild: part[0] == ':' || part[0] == '*'}
		parent.children = append(parent.children, child)
		parent = child
	}
	parent.fullPath = path
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.path == part {
			return child
		}
	}
	return nil
}

func (n *node) search(path string) *node {
	parent := n
	// 全路径相等的情况
	if parent.fullPath == path || parent.wildChild {
		return parent
	}
	for _, child := range parent.children {
		result := child.search(path)
		parent = child
		if result != nil {
			return result
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	parent := n
	nodes := make([]*node, 0)
	for _, child := range parent.children {
		if child.path == part || child.wildChild {
			nodes = append(nodes, child)
		}
		parent = child
	}
	return nodes
}
