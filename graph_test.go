package grapho

import (
	"testing"
)

func EqualsIntSlice(x, y []uint64) bool {
	if len(x) != len(y) {
		return false
	}

	for i, v := range x {
		if v != y[i] {
			return false
		}
	}
	return true
}

func TestGraphAddNode(t *testing.T) {
	g := NewGraph(false)

	g.AddNode(1, nil)

	attr := NewAttr()
	attr.Set("x", 25)
	g.AddNode(2, attr)

	if len(g.Nodes()) != 2 {
		t.Errorf("Expected size %d, got %d", 2, len(g.Nodes()))
	}

	attr, _ = g.Node(2)
	x, ok := attr.Get("x")
	if !ok || x.(int) != 25 {
		t.Errorf("Expected value: 25. Got %v", x)
	}

	// Update node
	attr = NewAttr()
	attr.Set("name", "Dylan")
	g.AddNode(1, attr)

	attr, _ = g.Node(1)
	name, ok := attr.Get("name")
	if !ok || name.(string) != "Dylan" {
		t.Errorf("Expected value: 'Dylan'. Got %v", name)
	}
}

func TestGraphDeleteNode(t *testing.T) {
	g := NewGraph(false)
	g.AddNode(1, nil)
	g.AddNode(2, nil)
	g.AddEdge(2, 1, 1, nil)

	g.DeleteNode(1)
	if len(g.Nodes()) != 1 {
		t.Errorf("Node was not successfully deleted")
	}
}

func TestGraphAddEdge(t *testing.T) {
	g := NewGraph(false)
	attr := NewAttr()
	attr.Set("x", 5)
	g.AddEdge(2, 1, 1, attr)

	nodes, ok := g.Neighbors(1)
	if !ok || !EqualsIntSlice(nodes, []uint64{2}) {
		t.Errorf("Edge was not successfully added")
	}
	nodes, ok = g.Neighbors(2)
	if !ok || !EqualsIntSlice(nodes, []uint64{1}) {
		t.Errorf("Edge was not successfully added")
	}

	edge, ok := g.Edge(1, 2)
	if !ok {
		t.Errorf("Edge was not successfully added")
	}
	x, ok := edge.Attr.Get("x")
	if !ok || x.(int) != 5 {
		t.Errorf("Expected value: 5. Got %v", x)
	}
}

func TestGraphDeleteEdge(t *testing.T) {
	g := NewGraph(false)
	g.AddEdge(2, 1, 1, NewAttr())
	g.DeleteEdge(1, 2)

	if _, ok := g.Edge(2, 1); ok {
		t.Errorf("Edge was not successfully deleted")
	}
}
