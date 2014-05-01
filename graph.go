package grapho

import (
	"errors"
	"fmt"
)

type Graph struct {
	// Each node has a slice of edges
	nodes map[uint64][]Edge
}

// NewGraph returns an empty Graph
func NewGraph() *Graph {
	return &Graph{make(map[uint64][]Edge)}
}

// Nodes returns the list of node ids (unsorted)
func (g *Graph) Nodes() []uint64 {
	nodes := make([]uint64, len(g.nodes))
	n := 0
	for k := range g.nodes {
		nodes[n] = k
		n++
	}
	return nodes
}

// AddNode adds the given node to the Graph. If the nodeId
// already exists, it will return an error
func (g *Graph) AddNode(nodeId uint64) error {
	if _, ok := g.nodes[nodeId]; ok {
		return fmt.Errorf("Node %d already exists", nodeId)
	}

	g.nodes[nodeId] = []Edge{}
	return nil
}

// AddEdge adds an Edge to the given node. If either the source
// or destiny node don't exist, an error is returned
func (g *Graph) AddEdge(nodeId uint64, edge Edge) error {
	edges, ok := g.nodes[nodeId]
	if !ok {
		return fmt.Errorf("Node %d doesn't exists", nodeId)
	} else if _, ok := g.nodes[edge.NodeId]; !ok {
		return fmt.Errorf("Node %d doesn't exists", edge.NodeId)
	}

	g.nodes[nodeId] = append(edges, edge)
	return nil
}

// GetEdges returns all the Edges associated with a particular node.
func (g *Graph) Edges(nodeId uint64) ([]Edge, error) {
	if vertices, ok := g.nodes[nodeId]; ok {
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
