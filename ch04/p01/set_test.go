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
				return p01.NewNode("")
			},
			want: ":[]",
		},
		{
			name: "key but no children",
			nodeCtor: func() *p01.Node {
				return p01.NewNode("key")
			},
			want: "key:[]",
		},
		{
			name: "key and a child",
			nodeCtor: func() *p01.Node {
				child := p01.NewNode("child")
				return p01.NewNode("parent", child)
			},
			want: "parent:[child]",
		},
		{
			name: "key and two children",
			nodeCtor: func() *p01.Node {
				c0 := p01.NewNode("child0")
				c1 := p01.NewNode("child1")
				return p01.NewNode("parent", c0, c1)
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
		"Has returns true after Adding":     sHasAfterAdding,
		"Get returns node after Adding":     sGetAfterAdding,
		"Has returns false after Deleteing": sNotHasAfterDeleting,
		"cannot Get after Deleteing":        sNotHasAfterDeleting,
		"string":                            sString,
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
		nodes []*p01.Node
	}{
		{nodes: nil},
		{nodes: []*p01.Node{n0}},
		{nodes: []*p01.Node{n0, n1}},
	}

	for _, test := range subtests {
		test := test
		name := fmt.Sprintf("%s", test.nodes)
		t.Run(name, func(t *testing.T) {
			s := p01.NewSet()

			for _, n := range test.nodes {
				s.Add(n)
			}

			for _, n := range test.nodes {
				if !s.Has(n.Key) {
					t.Errorf("missing node %s", n.Key)
				}
			}
		})
	}
}

func sGetAfterAdding(t *testing.T) {
	n0 := &p01.Node{Key: "0"}
	n1 := &p01.Node{Key: "1"}

	subtests := []struct {
		nodes []*p01.Node
	}{
		{nodes: nil},
		{nodes: []*p01.Node{n0}},
		{nodes: []*p01.Node{n0, n1}},
	}

	for _, test := range subtests {
		test := test
		name := fmt.Sprintf("%s", test.nodes)
		t.Run(name, func(t *testing.T) {
			s := p01.NewSet()

			for _, n := range test.nodes {
				s.Add(n)
			}

			for _, n := range test.nodes {
				got, ok := s.Get(n.Key)
				if !ok {
					t.Fatalf("missing node %s", n.Key)
				}

				if got != n {
					t.Errorf("\nwant %#v\n got %#v", n, got)
				}
			}
		})
	}
}

func sNotHasAfterDeleting(t *testing.T) {
	n0 := &p01.Node{Key: "0"}
	n1 := &p01.Node{Key: "1"}

	subtests := []struct {
		nodes []*p01.Node
	}{
		{nodes: []*p01.Node{n0}},
		{nodes: []*p01.Node{n0, n1}},
	}

	for _, test := range subtests {
		test := test
		name := fmt.Sprintf("%s", test.nodes)
		t.Run(name, func(t *testing.T) {
			s := p01.NewSet(test.nodes...)

			for _, n := range test.nodes {
				s.Delete(n.Key)
				if s.Has(n.Key) {
					t.Errorf("found node %s after deleting it", n.Key)
				}
			}
		})
	}
}

func sCannotGetAfterDeleting(t *testing.T) {
	n0 := &p01.Node{Key: "0"}
	n1 := &p01.Node{Key: "1"}

	subtests := []struct {
		nodes []*p01.Node
	}{
		{nodes: []*p01.Node{n0}},
		{nodes: []*p01.Node{n0, n1}},
	}

	for _, test := range subtests {
		test := test
		name := fmt.Sprintf("%s", test.nodes)
		t.Run(name, func(t *testing.T) {
			s := p01.NewSet(test.nodes...)

			for _, n := range test.nodes {
				s.Delete(n.Key)
				got, ok := s.Get(n.Key)
				if ok {
					t.Errorf("found node %s after deleting it", got.Key)
				}
			}
		})
	}
}

func sString(t *testing.T) {
	n0 := p01.NewNode("0")
	n1 := p01.NewNode("1")
	n0.Children.Add(n0)
	n0.Children.Add(n1)
	n1.Children.Add(n0)
	n1.Children.Add(n1)

	subtests := []struct {
		nodes []*p01.Node
		want  string
	}{
		{
			nodes: nil,
			want:  "[]",
		},
		{
			nodes: []*p01.Node{n0},
			want:  "[0]",
		},
		{
			nodes: []*p01.Node{n0, n1},
			want:  "[0 1]",
		},
	}

	for _, test := range subtests {
		test := test
		name := fmt.Sprintf("%s", test.nodes)
		t.Run(name, func(t *testing.T) {
			s := p01.NewSet(test.nodes...)
			got := s.String()
			if got != test.want {
				t.Errorf("\nwant %s\n got %s", test.want, got)
			}
		})
	}
}
