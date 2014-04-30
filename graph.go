package grapho

import "errors"

type Graph struct {
	// Each node has a slice of edges
	Nodes map[uint64][]Edge
}

// TODO: Add convenience methods to modify the Graph (AddNode, AddEdge, ...)

// GetEdges returns all the Edges associated with a particular node.
func (g Graph) GetEdges(nodeId uint64) ([]Edge, error) {
	if vertices, ok := g.Nodes[nodeId]; ok {
		return vertices, nil
	}
	return nil, errors.New("Node not found")
}

// Edge represents a relationship for a particular node
// It contains the NodeId of the successor and the cost to reach it
type Edge struct {
	NodeId uint64
	Cost   int
}
