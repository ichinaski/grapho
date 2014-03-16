package dijkstra

import "testing"

/**
 * struct that implements Graph interface.
 * state is represented with int values, holding the Node IDs
 */
type MyGraph struct {
    // Each node has a map of vertices, with their relative move cost
    nodes map[int]map[int]int
}

func (g MyGraph) GetChildren(nodeId NodeId) map[NodeId]int {
    children := make(map[NodeId]int)

    if vertices, ok := g.nodes[nodeId.(int)]; ok {
        for vertex, cost := range vertices {
            children[vertex] = cost
        }
    }

    return children
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
 * Test FindPath with a simple graph containing 5 vertices.
 * TODO: Read the graph from a file, allowing much larger graphs
 * to be tested.
 */
func TestFindPath(t *testing.T) {
    g := MyGraph {
        nodes: map[int]map[int]int {
            1: map[int]int { 2: 1, 3: 1, 5: 4},
            2: map[int]int { 4: 1 },
            3: map[int]int { 4: 1 },
            4: map[int]int { 5: 1 },
        },
    }

    path := FindPath(g, 1, 5)
    expected := []NodeId{1, 2, 4, 5}

    if !equalPath(path, expected) {
        t.Errorf("Path: %v. Expected: %v", path, expected)
    }

    path = FindPath(g, 1, 6)
    expected = []NodeId{}

    if !equalPath(path, expected) {
        t.Errorf("Path: %v. Expected: %v", path, expected)
    }
}
