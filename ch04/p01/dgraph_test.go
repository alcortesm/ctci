package p01_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/alcortesm/ctci/ch04/p01"
	"github.com/google/go-cmp/cmp"
)

func TestDGraph(t *testing.T) {
	t.Parallel()

	subtests := map[string]func(t *testing.T){
		"dump":                            dgDump,
		"hasRoute returns correct values": dgHasRouteOK,
		"hasRoute returns error if nodes are not found": dgHasRouteError,
	}

	for name, fn := range subtests {
		fn := fn
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			fn(t)
		})
	}
}

func dgDump(t *testing.T) {
	subtests := []p01.AdjacencyList{
		{},
		{"a": []string{}},
		{"a": []string{"a"}},
		{
			"a": []string{"b"},
			"b": []string{"a"},
		},
		{
			"a": []string{"a", "b"},
			"b": []string{"a", "b"},
		},
	}

	sortChildren := func(in []string) []string {
		out := append([]string(nil), in...)
		sort.Strings(out)
		return out
	}

	ignoreChildrenOrder := cmp.Transformer("Sort", sortChildren)

	for _, adjacencyList := range subtests {
		adjacencyList := adjacencyList
		name := fmt.Sprintf("%s", adjacencyList)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			g, err := p01.NewDGraph(adjacencyList)
			if err != nil {
				t.Fatal(err)
			}

			got := g.Dump()

			diff := cmp.Diff(adjacencyList, got, ignoreChildrenOrder)
			if diff != "" {
				t.Errorf("(-want +got)\n%s", diff)
			}
		})
	}
}

func dgHasRouteOK(t *testing.T) {
	type check struct {
		src  string
		dst  string
		want bool
	}

	subtests := []struct {
		scenario      string
		adjacencyList p01.AdjacencyList
		checks        []check
	}{
		{
			scenario:      "a", // just a node a
			adjacencyList: p01.AdjacencyList{"a": []string{}},
			checks: []check{
				{src: "a", dst: "a", want: false},
			},
		},
		{
			scenario:      "(a)", // a connected to itself
			adjacencyList: p01.AdjacencyList{"a": []string{"a"}},
			checks: []check{
				{src: "a", dst: "a", want: true},
			},
		},
		{
			scenario: "a   b", // two isolated nodes a and b
			adjacencyList: p01.AdjacencyList{
				"a": []string{},
				"b": []string{},
			},
			checks: []check{
				{src: "a", dst: "a", want: false},
				{src: "a", dst: "b", want: false},
				{src: "b", dst: "a", want: false},
				{src: "b", dst: "b", want: false},
			},
		},
		{
			scenario: "(a)   (b)", // two isolated nodes a and b, connected to themselves
			adjacencyList: p01.AdjacencyList{
				"a": []string{"a"},
				"b": []string{"b"},
			},
			checks: []check{
				{src: "a", dst: "a", want: true},
				{src: "a", dst: "b", want: false},
				{src: "b", dst: "a", want: false},
				{src: "b", dst: "b", want: true},
			},
		},
		{
			scenario: "a->b", // a connected to b
			adjacencyList: p01.AdjacencyList{
				"a": []string{"b"},
				"b": []string{},
			},
			checks: []check{
				{src: "a", dst: "a", want: false},
				{src: "a", dst: "b", want: true},
				{src: "b", dst: "a", want: false},
				{src: "b", dst: "b", want: false},
			},
		},
		{
			scenario: "(a)->b", // a connected to b and to itself
			adjacencyList: p01.AdjacencyList{
				"a": []string{"a", "b"},
				"b": []string{},
			},
			checks: []check{
				{src: "a", dst: "a", want: true},
				{src: "a", dst: "b", want: true},
				{src: "b", dst: "a", want: false},
				{src: "b", dst: "b", want: false},
			},
		},
		{
			scenario: "a->(b)", // a connected to b and to itself
			adjacencyList: p01.AdjacencyList{
				"a": []string{"b"},
				"b": []string{"b"},
			},
			checks: []check{
				{src: "a", dst: "a", want: false},
				{src: "a", dst: "b", want: true},
				{src: "b", dst: "a", want: false},
				{src: "b", dst: "b", want: true},
			},
		},
		{
			scenario: "(a)->(b)",
			adjacencyList: p01.AdjacencyList{
				"a": []string{"a", "b"},
				"b": []string{"a", "b"},
			},
			checks: []check{
				{src: "a", dst: "a", want: true},
				{src: "a", dst: "b", want: true},
				{src: "b", dst: "a", want: true},
				{src: "b", dst: "b", want: true},
			},
		},
		{
			scenario: "a<->b", // a connected to b and b connected to a
			adjacencyList: p01.AdjacencyList{
				"a": []string{"b"},
				"b": []string{"a"},
			},
			checks: []check{
				{src: "a", dst: "a", want: true},
				{src: "a", dst: "b", want: true},
				{src: "b", dst: "a", want: true},
				{src: "b", dst: "b", want: true},
			},
		},
		{
			scenario: "a<->(b)",
			adjacencyList: p01.AdjacencyList{
				"a": []string{"b"},
				"b": []string{"a", "b"},
			},
			checks: []check{
				{src: "a", dst: "a", want: true},
				{src: "a", dst: "b", want: true},
				{src: "b", dst: "a", want: true},
				{src: "b", dst: "b", want: true},
			},
		},
		{
			scenario: "a->b->c",
			adjacencyList: p01.AdjacencyList{
				"a": []string{"b"},
				"b": []string{"c"},
				"c": []string{},
			},
			checks: []check{
				{src: "a", dst: "a", want: false},
				{src: "a", dst: "b", want: true},
				{src: "a", dst: "c", want: true},
				{src: "b", dst: "a", want: false},
				{src: "b", dst: "b", want: false},
				{src: "b", dst: "c", want: true},
				{src: "c", dst: "a", want: false},
				{src: "c", dst: "b", want: false},
				{src: "c", dst: "c", want: false},
			},
		},
		{
			scenario: "a->b->c->a",
			adjacencyList: p01.AdjacencyList{
				"a": []string{"b"},
				"b": []string{"c"},
				"c": []string{"a"},
			},
			checks: []check{
				{src: "a", dst: "a", want: true},
				{src: "a", dst: "b", want: true},
				{src: "a", dst: "c", want: true},
				{src: "b", dst: "a", want: true},
				{src: "b", dst: "b", want: true},
				{src: "b", dst: "c", want: true},
				{src: "c", dst: "a", want: true},
				{src: "c", dst: "b", want: true},
				{src: "c", dst: "c", want: true},
			},
		},
		{
			scenario: "a->b  (c)",
			adjacencyList: p01.AdjacencyList{
				"a": []string{"b"},
				"b": []string{},
				"c": []string{"c"},
			},
			checks: []check{
				{src: "a", dst: "a", want: false},
				{src: "a", dst: "b", want: true},
				{src: "a", dst: "c", want: false},
				{src: "b", dst: "a", want: false},
				{src: "b", dst: "b", want: false},
				{src: "b", dst: "c", want: false},
				{src: "c", dst: "a", want: false},
				{src: "c", dst: "b", want: false},
				{src: "c", dst: "c", want: true},
			},
		},
		{
			scenario: "a<->b a->(c)",
			adjacencyList: p01.AdjacencyList{
				"a": []string{"b", "c"},
				"b": []string{"a", "b"},
				"c": []string{"c"},
			},
			checks: []check{
				{src: "a", dst: "a", want: true},
				{src: "a", dst: "b", want: true},
				{src: "a", dst: "c", want: true},
				{src: "b", dst: "a", want: true},
				{src: "b", dst: "b", want: true},
				{src: "b", dst: "c", want: true},
				{src: "c", dst: "a", want: false},
				{src: "c", dst: "b", want: false},
				{src: "c", dst: "c", want: true},
			},
		},
	}

	for _, test := range subtests {
		test := test

		for _, check := range test.checks {
			check := check
			name := fmt.Sprintf("%s|%s|%s",
				test.scenario, check.src, check.dst)

			t.Run(name, func(t *testing.T) {
				t.Parallel()

				g, err := p01.NewDGraph(test.adjacencyList)
				if err != nil {
					t.Fatalf("creating graph: %v", err)
				}

				got, err := g.HasRoute(check.src, check.dst)
				if err != nil {
					t.Fatalf("checking if there is a route: %v", err)
				}

				if got != check.want {
					t.Errorf("want %t, got %t", check.want, got)
				}
			})
		}
	}
}

func dgHasRouteError(t *testing.T) {
	type check struct {
		src string
		dst string
	}

	subtests := []struct {
		scenario      string
		adjacencyList p01.AdjacencyList
		checks        []check
	}{
		{
			scenario:      "empty",
			adjacencyList: nil,
			checks: []check{
				{src: "a", dst: "b"},
			},
		},
		{
			scenario: "a->b  (c)",
			adjacencyList: p01.AdjacencyList{
				"a": []string{"b"},
				"b": []string{},
				"c": []string{"c"},
			},
			checks: []check{
				{src: "a", dst: "d"},
				{src: "d", dst: "a"},
				{src: "d", dst: "d"},
			},
		},
	}

	for _, test := range subtests {
		test := test

		for _, check := range test.checks {
			check := check
			name := fmt.Sprintf("%s|%s|%s",
				test.scenario, check.src, check.dst)

			t.Run(name, func(t *testing.T) {
				t.Parallel()

				g, err := p01.NewDGraph(test.adjacencyList)
				if err != nil {
					t.Fatalf("creating graph: %v", err)
				}

				got, err := g.HasRoute(check.src, check.dst)
				if err == nil {
					t.Fatalf("unexpected success, got %t", got)
				}
			})
		}
	}
}
