package p01

import (
	"fmt"
)

// DGraph represents a directed graph.
type DGraph struct {
	Set
}

// AdjacencyList represents a DGraph as a collection of all its node
// keys, each with its list of adjacent nodes.
type AdjacencyList map[string][]string

// NewDGraph creates a DGraph from a adjacency list.
func NewDGraph(adjacencyList AdjacencyList) (*DGraph, error) {
	result := new(DGraph)
	result.Set = make(Set, len(adjacencyList))

	// two phases:
	//
	// Phase 1: create nodes for all the keys in the adjacencyList and
	// store them in result.nodes. They wont have any children yet.
	for nodeKey, _ := range adjacencyList {
		result.Add(&Node{
			Key: nodeKey,
		})
	}

	// Phase 2: go again over every key in the adjacencyList, adding
	// children to each node.
	for parentKey, childrenKeys := range adjacencyList {
		parent, _ := result.Get(parentKey)
		parent.Children = make(Set, len(childrenKeys))

		for _, k := range childrenKeys {
			children, ok := result.Get(k)
			if !ok {
				return nil, fmt.Errorf("unknown children %s in node %s",
					k, parentKey)
			}

			parent.Children.Add(children)
		}
	}

	return result, nil
}

// Dump returns the AdjacencyList for a graph.
func (g *DGraph) Dump() AdjacencyList {
	result := make(map[string][]string, len(g.Set))

	for key, node := range g.Set {
		children := make([]string, 0, len(node.Children))

		for k, _ := range node.Children {
			children = append(children, k)
		}

		result[key] = children
	}

	return result
}

// HasRoute returns whether there is a route from the node with key a to
// the node with key b.
func (g *DGraph) HasRoute(a, b string) (bool, error) {
	src, ok := g.Get(a)
	if !ok {
		return false, fmt.Errorf("missing node with key %s", a)
	}

	dst, ok := g.Get(b)
	if !ok {
		return false, fmt.Errorf("missing node with key %s", b)
	}

	// Breadth-first search
	queue := new(Queue)
	seen := make(Set) // protects us from enqueueing the same node twice

	queue.Enqueue(src)
	seen.Add(src)

	for !queue.IsEmpty() {
		current := queue.Dequeue().(*Node)

		for _, c := range current.Children {
			if c == dst {
				return true, nil
			}

			if !seen.Has(c.Key) {
				queue.Enqueue(c)
				seen.Add(c)
			}
		}
	}

	return false, nil
}
