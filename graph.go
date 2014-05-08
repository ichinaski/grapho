package grapho

import "sort"

// uint64Slice attaches the methods of sort.Interface to []uint64, sorting in increasing order.
type uint64Slice []uint64

func (p uint64Slice) Len() int           { return len(p) }
func (p uint64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p uint64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Edge represents a relationship between two nodes.
type Edge struct {
	Weight uint64 // Edge weight (cost)
	Attr   Attr   // Edge attribute set
}

func NewEdge(weight uint64, attr Attr) *Edge {
	if attr == nil {
		attr = NewAttr()
	}
	return &Edge{weight, attr}
}

// Graph represents an Undirected Graph.
// Each pair of nodes can only hold one edge between them.
type Graph struct {
	directed bool                        // true to represent a Digraph, false for Undirected Graphs
	nodes    map[uint64]Attr             // Nodes present in the Graph, with their attributes
	edges    map[uint64]map[uint64]*Edge // Adjacency list of edges, with their attributes
}

// NewGraph creates an empty Graph.
func NewGraph(directed bool) *Graph {
	return &Graph{
		directed: directed,
		nodes:    make(map[uint64]Attr),
		edges:    make(map[uint64]map[uint64]*Edge),
	}
}

// AddNode adds the given node to the Graph. If the node
// already exists, it will override its attributes.
func (g *Graph) AddNode(node uint64, attr Attr) {
	if attr == nil {
		attr = NewAttr()
	}

	g.nodes[node] = attr
	g.edges[node] = make(map[uint64]*Edge)
}

// DeleteNode removes a node entry from the Graph.
// Any edge associated with it will be removed too.
func (g *Graph) DeleteNode(node uint64) {
	// Remove incident edges
	for k := range g.edges[node] {
		delete(g.edges[k], node)
	}

	delete(g.edges, node)
	delete(g.nodes, node)
}

// AddEdge adds an edge (with its attributes) between nodes u and v
// If the nodes don't exist, they will be automatically created.
// If an u-v edge already existed, its attributes will be overridden.
func (g *Graph) AddEdge(u, v, weight uint64, attr Attr) {
	// Add nodes if necessary
	if _, ok := g.nodes[u]; !ok {
		g.AddNode(u, nil)
	}
	if _, ok := g.nodes[v]; !ok {
		g.AddNode(v, nil)
	}

	edge := NewEdge(weight, attr)

	g.edges[u][v] = edge
	if !g.directed {
		g.edges[v][u] = edge
	}
}

// DeleteEdge removes the u-v edge, if exists.
// If any of the nodes don't exist, nothing happens.
func (g *Graph) DeleteEdge(u, v uint64) {
	if _, ok := g.Node(u); ok {
		if _, ok := g.Node(v); ok {
			delete(g.edges[u], v)
			if !g.directed {
				delete(g.edges[v], u)
			}
		}
	}
}

// Nodes returns the list of nodes in the Graph (unsorted).
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
// a bool flag set to true if the node was found, false otherwise.
func (g *Graph) Node(node uint64) (Attr, bool) {
	attr, ok := g.nodes[node]
	return attr, ok
}

// Neighbors returns the list of nodes containing edges between the
// given node and them, ordered by ascending node uint64 value
// An extra bool flag determines whether the node was found.
func (g *Graph) Neighbors(node uint64) ([]uint64, bool) {
	if edges, ok := g.edges[node]; ok {
		nodes := make([]uint64, len(edges))

		n := 0
		for k := range edges {
			nodes[n] = k
			n++
		}
		sort.Sort(uint64Slice(nodes)) // order by node uint64 value

		return nodes, true
	}
	return nil, false
}

// Edge returns the Edge associated with the u-v node pair.
// An extra bool flag determines whether the edge was found.
// In undirected graphs, the edge u-v is be the same as v-u.
func (g *Graph) Edge(u, v uint64) (*Edge, bool) {
	if _, ok := g.Node(u); ok {
		if _, ok := g.Node(v); ok {
			edge, ok := g.edges[u][v]
			return edge, ok
		}
	}
	return nil, false
}
