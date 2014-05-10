package grapho

import (
	"testing"
)

func testEdgeExists(t *testing.T, g *Graph, u, v uint64, exist bool) {
	if _, ok := g.Edge(u, v); exist != ok {
		if exist {
			t.Errorf("Edge %d-%d should be present", u, v)
		} else {
			t.Errorf("Edge %d-%d should not be present", u, v)
		}
	}
}

func TestMinimumSpanningTree(t *testing.T) {
	g := NewGraph(false)

	g.AddEdge(1, 2, 1, nil)
	g.AddEdge(1, 3, 4, nil)
	g.AddEdge(1, 4, 3, nil)

	g.AddEdge(2, 4, 2, nil)
	g.AddEdge(3, 4, 5, nil)

	mst, err := MinimumSpanningTree(g, Prim)
	if err != nil {
		t.Fatalf("MinimumSpanningTree: %v", err)
	}
	if len(mst.Nodes()) != len(g.Nodes()) {
		t.Errorf("MST length: %d. Expected 4", len(mst.Nodes()))
	}

	// MST should have the edges 1-2, 1-3, 2-4
	testEdgeExists(t, mst, 1, 2, true)
	testEdgeExists(t, mst, 1, 3, true)
	testEdgeExists(t, mst, 2, 4, true)
	// Edges 1-4 and 3-4 should not be present
	testEdgeExists(t, mst, 1, 4, false)
	testEdgeExists(t, mst, 3, 4, false)

	g = NewGraph(false)

	g.AddEdge(1, 2, 2, nil)
	g.AddEdge(1, 4, 4, nil)

	g.AddEdge(2, 3, 2, nil)
	g.AddEdge(2, 4, 1, nil)
	g.AddEdge(2, 6, 4, nil)

	g.AddEdge(3, 6, 1, nil)
	g.AddEdge(3, 4, 3, nil)
	g.AddEdge(3, 5, 3, nil)

	g.AddEdge(4, 5, 2, nil)
	g.AddEdge(4, 6, 2, nil)

	g.AddEdge(5, 6, 1, nil)

	mst, err = MinimumSpanningTree(g, Prim)
	if err != nil {
		t.Fatalf("MinimumSpanningTree: %v", err)
	}
	if len(mst.Nodes()) != len(g.Nodes()) {
		t.Errorf("MST length: %d. Expected 4", len(mst.Nodes()))
	}

	// MST should have the edges 1-2, 2-3, 2-4, 3-6, 5-6
	testEdgeExists(t, mst, 1, 2, true)
	testEdgeExists(t, mst, 2, 3, true)
	testEdgeExists(t, mst, 2, 4, true)
	testEdgeExists(t, mst, 3, 6, true)
	testEdgeExists(t, mst, 5, 6, true)
	// Edges 1-4, 3-4, 4-5, 3-5, 2-6 and 4-6 should not be present
	testEdgeExists(t, mst, 1, 4, false)
	testEdgeExists(t, mst, 3, 4, false)
	testEdgeExists(t, mst, 4, 5, false)
	testEdgeExists(t, mst, 3, 5, false)
	testEdgeExists(t, mst, 2, 6, false)
	testEdgeExists(t, mst, 4, 6, false)
}
