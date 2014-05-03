package grapho

import (
	"testing"
)

// equalPath compares the given path with the expected path
func equalPath(s1, s2 []uint64) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

// TestDijkstra tests the Dijkstra algorithm with the sample test
func TestDijkstra(t *testing.T) {
	g := sampleDiGraph()

	expected := []uint64{1, 2, 5, 8}
	path, err := Dijkstra(g, 1, 8)
	if err != nil {
		t.Fatalf("Dijkstra: %v", err)
	} else if !equalPath(path, expected) {
		t.Errorf("Path: %v. Expected: %v", path, expected)
	}

	path, err = Dijkstra(g, 1, 10)
	if err == nil {
		t.Error("Dijkstra: Did not get expected error")
	}
}

// TestAstar tests the A* algorithm with an heuristic that
// assigns more priority to numbers closer to 9
func TestAStar(t *testing.T) {
	g := sampleDiGraph()

	h := func(node, goal uint64) int {
		return (int)(goal - node)
	}

	path, err := Astar(g, 1, 9, h)
	if err != nil {
		t.Fatalf("Astar: %v", err)
	}

	expected := []uint64{1, 3, 6, 8, 9}
	if !equalPath(path, expected) {
		t.Errorf("Path: %v. Expected: %v", path, expected)
	}
}

// TestBreathFirstSearch tests BreathFirstSearch algorithm with the sample test
func TestBreathFirstSearch(t *testing.T) {
	g := sampleDiGraph()

	path, err := BreathFirstSearch(g, 1, 9)
	if err != nil {
		t.Fatalf("BreathFirstSearch: %v", err)
	}

	expected := []uint64{1, 2, 4, 7, 9}
	if !equalPath(path, expected) {
		t.Errorf("Path: %v. Expected: %v", path, expected)
	}
}

// TestDepthFirstSearch tests DepthFirstSearch algorithm with the sample test
func TestDepthFirstSearch(t *testing.T) {
	g := sampleDiGraph()

	path, err := DepthFirstSearch(g, 1, 9)
	if err != nil {
		t.Fatalf("DepthFirstSearch: %v", err)
	}

	expected := []uint64{1, 2, 4, 7, 5, 3, 6, 8, 9}
	if !equalPath(path, expected) {
		t.Errorf("Path: %v. Expected: %v", path, expected)
	}
}

// sampleDiGraph creates a simple DiGraph for testing purposes
func sampleDiGraph() *DiGraph {
	//  1  2  4
	//  3  5  7
	//  6  8  9
	g := NewDiGraph()
	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)
	g.AddNode(4)
	g.AddNode(5)
	g.AddNode(6)
	g.AddNode(7)
	g.AddNode(8)
	g.AddNode(9)

	g.AddEdge(1, Edge{2, 1})
	g.AddEdge(1, Edge{3, 1})

	g.AddEdge(2, Edge{1, 1})
	g.AddEdge(2, Edge{4, 1})
	g.AddEdge(2, Edge{5, 1})

	g.AddEdge(3, Edge{1, 1})
	g.AddEdge(3, Edge{5, 1})
	g.AddEdge(3, Edge{6, 1})

	g.AddEdge(4, Edge{2, 1})
	g.AddEdge(4, Edge{7, 1})

	g.AddEdge(5, Edge{2, 1})
	g.AddEdge(5, Edge{3, 1})
	g.AddEdge(5, Edge{7, 1})
	g.AddEdge(5, Edge{8, 1})

	g.AddEdge(6, Edge{3, 1})
	g.AddEdge(6, Edge{8, 1})

	g.AddEdge(7, Edge{4, 1})
	g.AddEdge(7, Edge{5, 1})
	g.AddEdge(7, Edge{9, 1})

	g.AddEdge(8, Edge{5, 1})
	g.AddEdge(8, Edge{6, 1})
	g.AddEdge(8, Edge{9, 1})

	g.AddEdge(9, Edge{7, 1})
	g.AddEdge(9, Edge{8, 1})

	return g
}
