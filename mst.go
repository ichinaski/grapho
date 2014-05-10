// Minimum Spanning Tree
// Given a connected, undirected graph, calculate the minimum cost subgraph that connects all the vertices together.
// http://en.wikipedia.org/wiki/Minimum_spanning_tree

package grapho

import (
	"errors"
	"github.com/ichinaski/grapho/container"
)

type MstAlgorithm uint

const (
	Prim MstAlgorithm = iota
	// TODO: Kruskal
)

// Connected returns whether the Graph is fully connected or not.
func IsConnected(g *Graph) bool {
	// TODO: Implement connectivity check (Depth-first search?)
	return true
}

// mstState is the struct to be stored in the heap, holding a node and its parent
// The priority of the item is the edge weight between them
type mstState struct {
	node, parent uint64
}

// MinimumSpanningTree calculates the MST using the specified algorithm
func MinimumSpanningTree(graph *Graph, algo MstAlgorithm) (*Graph, error) {
	switch algo {
	case Prim:
		return PrimMst(graph)
	}

	return nil, errors.New("Unknown algorithm")
}

// PrimMst calculates the MST using PRIM algorithm (heap-based implementation).
func PrimMst(graph *Graph) (*Graph, error) {
	// Check if the Graph is undirected and connected
	if graph.IsDirected() {
		return nil, errors.New("Graph must be undirected")
	} else if !IsConnected(graph) {
		return nil, errors.New("Graph must be connected")
	}

	mst := NewGraph(false)
	pq := &container.PQueue{} // PQueue will determine which is the next node to add

	// expand adds the given node to the MST, and will recompute the PQueue priorities for its successors
	expand := func(node, parent uint64) {
		// Add node to the MST
		attr, _ := graph.Node(node)
		mst.AddNode(node, attr) // TODO: Deep copy of *Attr instead?

		if node != parent {
			// node == parent means that node has no parent.
			edge, _ := graph.Edge(node, parent)
			mst.AddEdge(node, parent, edge.Weight, edge.Attr)
		}

		// recompute priorities, if necessary
		succs, _ := graph.Neighbors(node)
		for _, succ := range succs {
			if _, ok := mst.Node(succ); !ok { // Skip the node if it's already in the mst Graph
				if edge, ok := graph.Edge(node, succ); ok {
					state := &mstState{succ, node}
					pq.Push(state, edge.Weight)
				}
			}
		}
	}

	node := graph.Nodes()[0] // choose randomly the first node
	expand(node, node)

	for pq.Len() > 0 { // NOTE: this invariant is not be the most efficient one. Compare the number of nodes instead
		state := pq.Pop().(*mstState)
		node, parent := state.node, state.parent

		// Only consider non expanded nodes (not present in mst)
		if _, ok := mst.Node(node); !ok {
			expand(node, parent)
		}
	}

	return mst, nil
}
