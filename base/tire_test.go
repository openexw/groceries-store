package base

import (
	"fmt"
	"testing"
)

func Test_node_insert(t *testing.T) {
	n := newNode()
	p := "/p/:id/create"
	parts := parsePattern(p)
	n.insert(p, parts, 0)

	p1 := "/p/:id/update"
	parts1 := parsePattern(p1)
	n.insert(p1, parts1, 0)

	p2 := "/p/:id/modify"
	parts2 := parsePattern(p2)
	n.insert(p2, parts2, 0)
}

func Test_node_search(t *testing.T) {
	n := newNode()
	p := "/p/:id/create"
	parts := parsePattern(p)
	n.insert(p, parts, 0)

	x := n.search(parts, 0)
	fmt.Printf("%v", x)
}
