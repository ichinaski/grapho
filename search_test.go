package search

import (
    "testing"
    "sort"
)

/**
 * struct that implements Graph interface.
 * state is represented with int values, holding the Node IDs
 */
type MyGraph struct {
    // Each node has a map of vertices, with their relative move cost
    nodes map[int]map[int]int
}

func (g MyGraph) GetEdges(nodeId NodeId) []Edge {
    var edges []Edge

    if vertices, ok := g.nodes[nodeId.(int)]; ok {
        // Sort vertices, to queue them ordered in the slice
        var keys []int
        for k := range vertices {
            keys = append(keys, k)
        }
        sort.Ints(keys)
        for _, k := range keys {
            edges = append(edges, Edge { k, vertices[k] })
        }
    }

    return edges
}

// equalPath compares the given path with the expected path
func equalPath(s1, s2 []NodeId) bool {
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

/**
 * Test Dijkstra with a simple graph containing 5 vertices.
 * TODO: Read the graph from a file, allowing much larger graphs
 * to be tested.
 */
func TestDijkstra(t *testing.T) {
    g := MyGraph {
        nodes: map[int]map[int]int {
            1: map[int]int { 2: 1, 3: 1, 5: 4},
            2: map[int]int { 4: 1 },
            3: map[int]int { 4: 1 },
            4: map[int]int { 5: 1 },
        },
    }

    path := Dijkstra(g, 1, 5)
    expected := []NodeId{1, 2, 4, 5}

    if !equalPath(path, expected) {
        t.Errorf("Path: %v. Expected: %v", path, expected)
    }

    path = Dijkstra(g, 1, 6)
    expected = []NodeId{}

    if !equalPath(path, expected) {
        t.Errorf("Path: %v. Expected: %v", path, expected)
    }
}

/**
 * Test Astar with a simple 3x3 grid and an heuristic that
 * assigns more priority to numbers closer to 9
 */
func TestAStar(t *testing.T) {
    //  1  2  4
    //  3  5  7
    //  6  8  9
    g := MyGraph {
        nodes: map[int]map[int]int {
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

    h := func(node, goal NodeId) int {
        return goal.(int) - node.(int)
    }

    path := Astar(g, 1, 9, h)
    expected := []NodeId{1, 3, 6, 8, 9}

    if !equalPath(path, expected) {
        t.Errorf("Path: %v. Expected: %v", path, expected)
    }
}

/**
 * Test BreathFirstSearch with a simple 3x3 grid
 */
func TestBreathFirstSearch(t *testing.T) {
    //  1  2  4
    //  3  5  7
    //  6  8  9
    g := MyGraph {
        nodes: map[int]map[int]int {
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
    expected := []NodeId{1, 2, 4, 7, 9}

    if !equalPath(path, expected) {
        t.Errorf("Path: %v. Expected: %v", path, expected)
    }
}

/**
 * Test DepthFirstSearch with a simple 3x3 grid
 */
func TestDepthFirstSearch(t *testing.T) {
    //  1  2  4
    //  3  5  7
    //  6  8  9
    g := MyGraph {
        nodes: map[int]map[int]int {
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
    expected := []NodeId{1, 2, 4, 7, 5, 3, 6, 8, 9}

    if !equalPath(path, expected) {
        t.Errorf("Path: %v. Expected: %v", path, expected)
    }
}
