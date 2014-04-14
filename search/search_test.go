package search

import (
    "github.com/ichinaski/graph-utils/graph"
    "testing"
)

// equalPath compares the given path with the expected path
func equalPath(s1, s2 []int) bool {
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
    g := graph.Graph {
        Nodes: map[int]map[int]int {
            1: map[int]int { 2: 1, 3: 1, 5: 4},
            2: map[int]int { 4: 1 },
            3: map[int]int { 4: 1 },
            4: map[int]int { 5: 1 },
        },
    }

    path := Dijkstra(g, 1, 5)
    expected := []int{1, 2, 4, 5}

    if !equalPath(path, expected) {
        t.Errorf("Path: %v. Expected: %v", path, expected)
    }

    path = Dijkstra(g, 1, 6)
    expected = []int{}

    if !equalPath(path, expected) {
        t.Errorf("Path: %v. Expected: %v", path, expected)
    }
}

// Test Astar with a simple 3x3 grid and an heuristic that
// assigns more priority to numbers closer to 9
func TestAStar(t *testing.T) {
    //  1  2  4
    //  3  5  7
    //  6  8  9
    g := graph.Graph {
        Nodes: map[int]map[int]int {
            1: map[int]int { 2: 1, 3: 1},
            2: map[int]int { 1: 1, 4:1, 5:1 },
            3: map[int]int { 1: 1, 5:1, 6:1 },
            4: map[int]int { 2: 1, 7:1 },
            5: map[int]int { 2: 1, 3:1, 7:1, 8:1 },
            6: map[int]int { 3: 1, 8:1 },
            7: map[int]int { 4: 1, 5:1, 9:1 },
            8: map[int]int { 5: 1, 6:1, 9:1 },
            9: map[int]int { 7: 1, 8:1 },
        },
    }

    h := func(node, goal int) int {
        return goal - node
    }

    path := Astar(g, 1, 9, h)
    expected := []int{1, 3, 6, 8, 9}

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
        Nodes: map[int]map[int]int {
            1: map[int]int { 2: 1, 3: 1},
            2: map[int]int { 1: 1, 4:1, 5:1 },
            3: map[int]int { 1: 1, 5:1, 6:1 },
            4: map[int]int { 2: 1, 7:1 },
            5: map[int]int { 2: 1, 3:1, 7:1, 8:1 },
            6: map[int]int { 3: 1, 8:1 },
            7: map[int]int { 4: 1, 5:1, 9:1 },
            8: map[int]int { 5: 1, 6:1, 9:1 },
            9: map[int]int { 7: 1, 8:1 },
        },
    }

    path := BreathFirstSearch(g, 1, 9)
    expected := []int{1, 2, 4, 7, 9}

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
        Nodes: map[int]map[int]int {
            1: map[int]int { 2: 1, 3: 1},
            2: map[int]int { 1: 1, 4:1, 5:1 },
            3: map[int]int { 1: 1, 5:1, 6:1 },
            4: map[int]int { 2: 1, 7:1 },
            5: map[int]int { 2: 1, 3:1, 7:1, 8:1 },
            6: map[int]int { 3: 1, 8:1 },
            7: map[int]int { 4: 1, 5:1, 9:1 },
            8: map[int]int { 5: 1, 6:1, 9:1 },
            9: map[int]int { 7: 1, 8:1 },
        },
    }

    path := DepthFirstSearch(g, 1, 9)
    expected := []int{1, 2, 4, 7, 5, 3, 6, 8, 9}

    if !equalPath(path, expected) {
        t.Errorf("Path: %v. Expected: %v", path, expected)
    }
}
