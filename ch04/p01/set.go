package p01

import (
	"fmt"
	"sort"
)

type Node struct {
	Key      string
	Children Set
}

func (n Node) String() string {
	return fmt.Sprintf("%s:%s", n.Key, n.Children)
}

type Set map[string]*Node

func (s Set) String() string {
	return fmt.Sprintf("%s", s.Keys())
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

// Keys returns the keys in the set sorted alphabetically.
func (s Set) Keys() []string {
	result := make([]string, 0, len(s))

	for k, _ := range s {
		result = append(result, k)
	}

	sort.Strings(result)

	return result
}
