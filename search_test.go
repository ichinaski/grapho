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
	g := sampleGraph()

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
	g := sampleGraph()

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
	g := sampleGraph()

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
	g := sampleGraph()

	path, err := DepthFirstSearch(g, 1, 9)
	if err != nil {
		t.Fatalf("DepthFirstSearch: %v", err)
	}

	expected := []uint64{1, 2, 4, 7, 5, 3, 6, 8, 9}
	if !equalPath(path, expected) {
		t.Errorf("Path: %v. Expected: %v", path, expected)
	}
}

// sampleGraph creates a simple Graph for testing purposes
func sampleGraph() *Graph {
	//  1  2  4
	//  3  5  7
	//  6  8  9
	g := NewGraph(true) // Directed Graph
	g.AddNode(1, nil)
	g.AddNode(2, nil)
	g.AddNode(3, nil)
	g.AddNode(4, nil)
	g.AddNode(5, nil)
	g.AddNode(6, nil)
	g.AddNode(7, nil)
	g.AddNode(8, nil)
	g.AddNode(9, nil)

	g.AddEdge(1, 2, 1, nil)
	g.AddEdge(1, 3, 1, nil)

	g.AddEdge(2, 1, 1, nil)
	g.AddEdge(2, 4, 1, nil)
	g.AddEdge(2, 5, 1, nil)

	g.AddEdge(3, 1, 1, nil)
	g.AddEdge(3, 5, 1, nil)
	g.AddEdge(3, 6, 1, nil)

	g.AddEdge(4, 2, 1, nil)
	g.AddEdge(4, 7, 1, nil)

	g.AddEdge(5, 2, 1, nil)
	g.AddEdge(5, 3, 1, nil)
	g.AddEdge(5, 7, 1, nil)
	g.AddEdge(5, 8, 1, nil)

	g.AddEdge(6, 3, 1, nil)
	g.AddEdge(6, 8, 1, nil)

	g.AddEdge(7, 4, 1, nil)
	g.AddEdge(7, 5, 1, nil)
	g.AddEdge(7, 9, 1, nil)

	g.AddEdge(8, 5, 1, nil)
	g.AddEdge(8, 6, 1, nil)
	g.AddEdge(8, 9, 1, nil)

	g.AddEdge(9, 7, 1, nil)
	g.AddEdge(9, 8, 1, nil)

	return g
}
