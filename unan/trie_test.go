package unan

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_newNode(t *testing.T) {
	trie := newNode()
	trie.insertChild("/user/create")
	trie.insertChild("/user/update")
	trie.insertChild("/user/:id")
}

func Test_node_matchChildren(t *testing.T) {
	type fields struct {
		path      string
		children  []*node
		wildChild bool
		fullPath  string
	}
	type args struct {
		part string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				path:      tt.fields.path,
				children:  tt.fields.children,
				wildChild: tt.fields.wildChild,
				fullPath:  tt.fields.fullPath,
			}
			if got := n.matchChildren(tt.args.part); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matchChildren() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_insertChild(t *testing.T) {
	n := newNode()
	//n.insertChild("user")
	n.insertChild("/user/create")
	//n.insertChild("/user/update")
	//n.insertChild("/user/:id")
	x := n.search("/user/create")
	fmt.Printf("%v", x)
}

func addChild(n *node, path []string) {
	for _, s := range path {
		n.insertChild(s)
	}
}
func Test_node_search(t *testing.T) {
	n := newNode()
	p1 := "/user/create"
	p2 := "/user/update"
	p3 := "/user/delete"
	p4 := "/user/:id/info"
	p5 := "/static/*filepath"
	addChild(n, []string{p1, p2, p3, p4, p5})

	r := n.search(p4)
	fmt.Printf("%v", r)
}
