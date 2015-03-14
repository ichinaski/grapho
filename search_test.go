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

// TestSearch tests all 4 search algorithms, with both directed and undirected graphs
func TestSearch(t *testing.T) {
	// Dijkstra
	testDijkstra(t, sampleGraph())
	testDijkstra(t, sampleDiGraph())

	// A*
	testAstar(t, sampleGraph())
	testAstar(t, sampleDiGraph())

	// BFS
	testBreadthFirstSearch(t, sampleGraph())
	testBreadthFirstSearch(t, sampleDiGraph())

	// DFS
	testDepthFirstSearch(t, sampleGraph())
	testDepthFirstSearch(t, sampleDiGraph())
}

// testDijkstra tests the Dijkstra algorithm with the given graph
func testDijkstra(t *testing.T, g *Graph) {
	expected := []uint64{1, 2, 5, 8}
	path, err := Search(g, 1, 8, Dijkstra, nil)
	if err != nil {
		t.Fatalf("Dijkstra: %v", err)
	} else if !equalPath(path, expected) {
		t.Errorf("Path: %v. Expected: %v", path, expected)
	}

	path, err = Search(g, 1, 10, Dijkstra, nil)
	if err == nil {
		t.Error("Dijkstra: Did not get expected error")
	}
}

// testAstar tests the A* algorithm with an heuristic that
// assigns more priority to numbers closer to 9
func testAstar(t *testing.T, g *Graph) {
	h := func(node, goal uint64) int {
		return (int)(goal - node)
	}

	path, err := Search(g, 1, 9, Astar, h)
	if err != nil {
		t.Fatalf("Astar: %v", err)
	}

	expected := []uint64{1, 3, 6, 8, 9}
	if !equalPath(path, expected) {
		t.Errorf("Path: %v. Expected: %v", path, expected)
	}
}

// testBreadthFirstSearch tests BreadthFirstSearch algorithm with the given graph
func testBreadthFirstSearch(t *testing.T, g *Graph) {
	path, err := Search(g, 1, 9, BreadthFirstSearch, nil)
	if err != nil {
		t.Fatalf("BreadthFirstSearch: %v", err)
	}

	expected := []uint64{1, 2, 4, 7, 9}
	if !equalPath(path, expected) {
		t.Errorf("Path: %v. Expected: %v", path, expected)
	}
}

// testDepthFirstSearch tests DepthFirstSearch algorithm with the given graph
func testDepthFirstSearch(t *testing.T, g *Graph) {
	path, err := Search(g, 1, 9, DepthFirstSearch, nil)
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
	g := NewGraph(false) // Directed Graph

	g.AddEdge(1, 2, 1, nil)
	g.AddEdge(1, 3, 1, nil)
	g.AddEdge(2, 4, 1, nil)
	g.AddEdge(2, 5, 1, nil)
	g.AddEdge(3, 5, 1, nil)
	g.AddEdge(3, 6, 1, nil)
	g.AddEdge(4, 7, 1, nil)
	g.AddEdge(5, 7, 1, nil)
	g.AddEdge(5, 8, 1, nil)
	g.AddEdge(6, 8, 1, nil)
	g.AddEdge(7, 9, 1, nil)
	g.AddEdge(8, 9, 1, nil)

	return g
}

// sampleDiGraph creates a simple DiGraph for testing purposes
func sampleDiGraph() *Graph {
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
