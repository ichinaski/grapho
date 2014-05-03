package grapho

import (
	"testing"
)

func TestNodes(t *testing.T) {
	g := NewDiGraph()
	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)

	nodes := g.Nodes()
	if len(nodes) != 3 {
		t.Errorf("Expected size %d, got %d", 3, len(nodes))
	}
}

func TestAddNode(t *testing.T) {
	graph := NewDiGraph()
	graph.AddNode(1)
	graph.AddNode(2)
	graph.AddNode(3)

	if err := graph.AddNode(1); err == nil {
		t.Errorf("Expected error (node exists) not returned")
	}

}

func TestAddEdge(t *testing.T) {
	g := NewDiGraph()
	g.AddNode(1)

	if err := g.AddEdge(1, Edge{2, 1}); err == nil {
		t.Errorf("Expected error (node doesn't exist) not returned")
	}

	g.AddNode(2)
	g.AddNode(3)

	if err := g.AddEdge(1, Edge{2, 1}); err != nil {
		t.Errorf("Error %v, err")
	}

	g.AddEdge(1, Edge{3, 1})

	edges, _ := g.Edges(1)
	if len(edges) != 2 {
		t.Errorf("Expected size %d, got %d", 2, len(edges))
	}
}
