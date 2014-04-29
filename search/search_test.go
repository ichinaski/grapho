package search

import (
    "github.com/ichinaski/graph-utils/graph"
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

// Test Dijkstra with a simple graph containing 5 vertices.
func TestDijkstra(t *testing.T) {
    //                  4
    //   ________________________________
    //  / 1        1        1       1    \
    // 1-----> 2 -----> 3 -----> 4 -----> 5
    //  \_______1_______/       
    g := graph.Graph {
        Nodes: map[uint64][]graph.Edge {
            1: []graph.Edge { {2, 1}, {3, 1}, {5, 4} },
            2: []graph.Edge { {4, 1} },
            3: []graph.Edge { {4, 1} },
            4: []graph.Edge { {5, 1} },
        },
    }

    expected := []uint64{1, 2, 4, 5}

    path, err := Dijkstra(g, 1, 5)
    if err != nil {
        t.Fatalf("Dijkstra: %v", err)
    }
    if !equalPath(path, expected) {
        t.Errorf("Path: %v. Expected: %v", path, expected)
    }

    path, err = Dijkstra(g, 1, 6)
    if err == nil {
        t.Error("Dijkstra: Did not get expected error")
    }
}

// Test Astar with a simple 3x3 grid and an heuristic that
// assigns more priority to numbers closer to 9
func TestAStar(t *testing.T) {
    //  1  2  4
    //  3  5  7
    //  6  8  9
    g := graph.Graph {
        Nodes: map[uint64][]graph.Edge {
            1: []graph.Edge { {2, 1}, {3, 1} },
            2: []graph.Edge { {1, 1}, {4, 1}, {5, 1} },
            3: []graph.Edge { {1, 1}, {5, 1}, {6, 1} },
            4: []graph.Edge { {2, 1}, {7, 1} },
            5: []graph.Edge { {2, 1}, {3, 1}, {7, 1}, {8, 1} },
            6: []graph.Edge { {3, 1}, {8, 1} },
            7: []graph.Edge { {4, 1}, {5, 1}, {9, 1} },
            8: []graph.Edge { {5, 1}, {6, 1}, {9, 1} },
            9: []graph.Edge { {7, 1}, {8, 1} },
        },
    }

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

// Test BreathFirstSearch with a simple 3x3 grid
func TestBreathFirstSearch(t *testing.T) {
    //  1  2  4
    //  3  5  7
    //  6  8  9
    g := graph.Graph {
        Nodes: map[uint64][]graph.Edge {
            1: []graph.Edge { {2, 1}, {3, 1} },
            2: []graph.Edge { {1, 1}, {4, 1}, {5, 1} },
            3: []graph.Edge { {1, 1}, {5, 1}, {6, 1} },
            4: []graph.Edge { {2, 1}, {7, 1} },
            5: []graph.Edge { {2, 1}, {3, 1}, {7, 1}, {8, 1} },
            6: []graph.Edge { {3, 1}, {8, 1} },
            7: []graph.Edge { {4, 1}, {5, 1}, {9, 1} },
            8: []graph.Edge { {5, 1}, {6, 1}, {9, 1} },
            9: []graph.Edge { {7, 1}, {8, 1} },
        },
    }

    path, err := BreathFirstSearch(g, 1, 9)
    if err != nil {
        t.Fatalf("BreathFirstSearch: %v", err)
    }

    expected := []uint64{1, 2, 4, 7, 9}
    if !equalPath(path, expected) {
        t.Errorf("Path: %v. Expected: %v", path, expected)
    }
}

// Test DepthFirstSearch with a simple 3x3 grid
func TestDepthFirstSearch(t *testing.T) {
    //  1  2  4
    //  3  5  7
    //  6  8  9
    g := graph.Graph {
        Nodes: map[uint64][]graph.Edge {
            1: []graph.Edge { {2, 1}, {3, 1} },
            2: []graph.Edge { {1, 1}, {4, 1}, {5, 1} },
            3: []graph.Edge { {1, 1}, {5, 1}, {6, 1} },
            4: []graph.Edge { {2, 1}, {7, 1} },
            5: []graph.Edge { {2, 1}, {3, 1}, {7, 1}, {8, 1} },
            6: []graph.Edge { {3, 1}, {8, 1} },
            7: []graph.Edge { {4, 1}, {5, 1}, {9, 1} },
            8: []graph.Edge { {5, 1}, {6, 1}, {9, 1} },
            9: []graph.Edge { {7, 1}, {8, 1} },
        },
    }

    path, err := DepthFirstSearch(g, 1, 9)
    if err != nil {
        t.Fatalf("DepthFirstSearch: %v", err)
    }

    expected := []uint64{1, 2, 4, 7, 5, 3, 6, 8, 9}
    if !equalPath(path, expected) {
        t.Errorf("Path: %v. Expected: %v", path, expected)
    }
}
