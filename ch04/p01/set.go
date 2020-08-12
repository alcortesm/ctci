package p01

import (
	"fmt"
	"sort"
)

// Node represents a node in a graph.
// It is not zero value safe, use NewNode instead.
type Node struct {
	Key      string
	Children Set
}

func NewNode(key string, children ...*Node) *Node {
	result := new(Node)
	result.Key = key
	result.Children = NewSet(children...)

	return result
}

func (n Node) String() string {
	return fmt.Sprintf("%s:%s", n.Key, n.Children)
}

type Set map[string]*Node

// NewSet creates a new set and Adds the nodes to it.
func NewSet(nodes ...*Node) Set {
	result := make(Set)

	for _, n := range nodes {
		result.Add(n)
	}

	return result
}

func (s Set) Has(key string) bool {
	_, ok := s[key]
	return ok
}

func (s Set) Get(key string) (*Node, bool) {
	n, ok := s[key]
	return n, ok
}

func (s Set) Add(n *Node) {
	s[n.Key] = n
}

func (s Set) Delete(key string) {
	delete(s, key)
}

// String returns the keys in the set sorted alphabetically.
func (s Set) String() string {
	result := make([]string, 0, len(s))

	for k, _ := range s {
		result = append(result, k)
	}

	sort.Strings(result)

	return fmt.Sprintf("%s", result)
}
