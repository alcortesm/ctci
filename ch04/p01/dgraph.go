package p01

import "fmt"

type DGraph struct {
	nodes Set
}

// NewDGraph creates a DGraph from a adjacency list.
func NewDGraph(adjacencyList map[string][]string) (*DGraph, error) {
	result := new(DGraph)
	result.nodes = make(Set, len(adjacencyList))

	// two phases:
	//
	// Phase 1: create nodes for all the keys in the adjacencyList and
	// store them in result.nodes. They wont have any children yet.
	for nodeKey, _ := range adjacencyList {
		result.nodes.Add(&Node{
			Key: nodeKey,
		})
	}

	// Phase 2: go again over every key in the adjacencyList, adding
	// children to each node.
	for parentKey, childrenKeys := range adjacencyList {
		parent, _ := result.nodes.Get(parentKey)
		parent.Children = make(Set, len(childrenKeys))

		for _, k := range childrenKeys {
			children, ok := result.nodes.Get(k)
			if !ok {
				return nil, fmt.Errorf("unknown children %s in node %s",
					k, parentKey)
			}

			parent.Children.Add(children)
		}
	}

	return result, nil
}

func (g *DGraph) Dump() map[string][]string {
	result := make(map[string][]string, len(g.nodes))

	for key, node := range g.nodes {
		result[key] = node.Children.Keys()
	}

	return result
}
