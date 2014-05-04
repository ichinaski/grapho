package grapho

import (
	"errors"
	"fmt"
)

// Attr is a set of attributes associated to a node/edge
// Keys are strings. Values can be anything.
type Attr struct {
	attr map[string]interface{}
}

// NewAttr creates an empty set of attributes
func NewAttr() *Attr {
	attr := Attr{make(map[string]interface{})}
	return &attr
}

// Set associates a new key-value pair
func (attr *Attr) Set(key string, value interface{}) {
	attr.attr[key] = value
}

// Get returns the associated value to the given key, and
// a bool set to true if the key was found, false otherwise
func (attr *Attr) Get(key string) (interface{}, bool) {
	v, ok := attr.attr[key]
	return v, ok
}

// Graph represents an Undirected Graph
// TODO: Create a common interface for Graph and DiGraph
type Graph struct {
	nodes map[uint64]*Attr            // Nodes present in the Graph, with their attributes
	edges map[uint64]map[uint64]*Attr // Adjacency list of edges, with their attributes
}

// NewGraph creates an empty Graph
func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[uint64]*Attr),
		edges: make(map[uint64]map[uint64]*Attr),
	}
}

// AddNode adds the given node to the Graph. If the nodeId
// already exists, it will override its attributes
func (g *Graph) AddNode(nodeId uint64, attr *Attr) {
	if attr == nil {
		attr = NewAttr()
	}

	g.nodes[nodeId] = attr
	g.edges[nodeId] = make(map[uint64]*Attr)
}

// DeleteNode removes a node entry from the Graph.
// Any edge associated with it will be removed too.
func (g *Graph) DeleteNode(nodeId uint64) {
	// Remove incident edges
	for k := range g.edges[nodeId] {
		delete(g.edges[k], nodeId)
	}

	delete(g.edges, nodeId)
	delete(g.nodes, nodeId)
}

// AddEdge adds an edge (with its attributes) between nodes u and v
// If the nodes don't exist, they will be automatically created.
// If an u-v edge already existed, its attributes will be overridden
func (g *Graph) AddEdge(u, v uint64, attr *Attr) {
	if attr == nil {
		attr = NewAttr()
	}

	// Add nodes if necessary
	if _, ok := g.nodes[u]; !ok {
		g.AddNode(u, nil)
	}
	if _, ok := g.nodes[v]; !ok {
		g.AddNode(v, nil)
	}

	g.edges[u][v] = attr
	g.edges[v][u] = attr
}

// DeleteEdge removes the u-v edge, if exists.
// If any of the nodes don't exist, nothing happens
func (g *Graph) DeleteEdge(u, v uint64) {
	if _, ok := g.Node(u); ok {
		if _, ok := g.Node(v); ok {
			delete(g.edges[u], v)
			delete(g.edges[v], u)
		}
	}
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

// Node returns the attributes associated with a given node, and
// a bool flag set to true if the node was found, false otherwise
func (g *Graph) Node(nodeId uint64) (*Attr, bool) {
	attr, ok := g.nodes[nodeId]
	return attr, ok
}

// Neighbors returns the list of nodes containing edges between the
// given node and them.
// An extra bool flag determines whether the node was found.
func (g *Graph) Neighbors(nodeId uint64) ([]uint64, bool) {
	if edges, ok := g.edges[nodeId]; ok {
		nodes := make([]uint64, len(edges))

		n := 0
		for k := range edges {
			nodes[n] = k
			n++
		}
		return nodes, true
	}
	return nil, false
}

// Edge return the attributes associated with the u-v edge.
// An extra bool flag determines whether the edge was found.
func (g *Graph) Edge(u, v uint64) (*Attr, bool) {
	if _, ok := g.Node(u); ok {
		if _, ok := g.Node(v); ok {
			edges, ok := g.edges[u][v]
			return edges, ok
		}
	}
	return nil, false
}

/*********************/
/****** DiGraph ******/
/*********************/

// DiGraph represents a Directed Graph
type DiGraph struct {
	// Each node has a slice of edges
	nodes map[uint64][]Edge
}

// NewDiGraph returns an empty DiGraph
func NewDiGraph() *DiGraph {
	return &DiGraph{make(map[uint64][]Edge)}
}

// Nodes returns the list of node ids (unsorted)
func (g *DiGraph) Nodes() []uint64 {
	nodes := make([]uint64, len(g.nodes))
	n := 0
	for k := range g.nodes {
		nodes[n] = k
		n++
	}
	return nodes
}

// AddNode adds the given node to the DiGraph. If the nodeId
// already exists, it will return an error
func (g *DiGraph) AddNode(nodeId uint64) error {
	if _, ok := g.nodes[nodeId]; ok {
		return fmt.Errorf("Node %d already exists", nodeId)
	}

	g.nodes[nodeId] = []Edge{}
	return nil
}

// AddEdge adds an Edge to the given node. If either the source
// or destiny node don't exist, an error is returned
func (g *DiGraph) AddEdge(nodeId uint64, edge Edge) error {
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
func (g *DiGraph) Edges(nodeId uint64) ([]Edge, error) {
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
