package p01_test

import (
	"fmt"
	"testing"

	"github.com/alcortesm/ctci/ch04/p01"
)

func TestNode_String(t *testing.T) {
	t.Parallel()

	subtests := []struct {
		name     string
		nodeCtor func() *p01.Node
		want     string
	}{
		{
			name: "empty key and no children",
			nodeCtor: func() *p01.Node {
				return &p01.Node{}
			},
			want: ":[]",
		},
		{
			name: "key but no children",
			nodeCtor: func() *p01.Node {
				return &p01.Node{
					Key: "key",
				}
			},
			want: "key:[]",
		},
		{
			name: "key and a child",
			nodeCtor: func() *p01.Node {
				child := &p01.Node{
					Key: "child",
				}

				parent := &p01.Node{
					Key:      "parent",
					Children: make(p01.Set),
				}

				parent.Children.Add(child)

				return parent
			},
			want: "parent:[child]",
		},
		{
			name: "key and two children",
			nodeCtor: func() *p01.Node {
				c0 := &p01.Node{
					Key: "child0",
				}

				c1 := &p01.Node{
					Key: "child1",
				}

				parent := &p01.Node{
					Key:      "parent",
					Children: make(p01.Set),
				}

				parent.Children.Add(c0)
				parent.Children.Add(c1)

				return parent
			},
			want: "parent:[child0 child1]",
		},
	}

	for _, test := range subtests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			node := test.nodeCtor()
			got := node.String()
			if got != test.want {
				t.Errorf("want %q\n got %q", test.want, got)
			}
		})
	}
}

func TestSet(t *testing.T) {
	t.Parallel()

	subtests := map[string]func(t *testing.T){
		"has returns true after adding": sHasAfterAdding,
	}

	for name, fn := range subtests {
		fn := fn
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			fn(t)
		})
	}
}

func sHasAfterAdding(t *testing.T) {
	n0 := &p01.Node{Key: "0"}
	n1 := &p01.Node{Key: "1"}

	subtests := []struct {
		toAdd []*p01.Node
	}{
		{toAdd: nil},
		{toAdd: []*p01.Node{n0}},
		{toAdd: []*p01.Node{n0, n1}},
	}

	for _, test := range subtests {
		test := test
		name := fmt.Sprintf("%s", test.toAdd)
		t.Run(name, func(t *testing.T) {
			s := make(p01.Set)
			for _, n := range test.toAdd {
				s.Add(n)
			}

			for _, n := range test.toAdd {
				if !s.Has(n.Key) {
					t.Errorf("missing node %s", n.Key)
				}
			}
		})
	}
}
