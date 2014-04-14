package graph

import "sort"

type Graph struct {
    // Each node has a map of vertices, with their relative move cost
    Nodes map[int]map[int]int
}

func (g Graph) GetEdges(nodeId int) []Edge {
    var edges []Edge

    if vertices, ok := g.Nodes[nodeId]; ok {
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


// Edge represents a relationship for a particular node
// It contains the NodeId of the successor and the cost to reach it
// Note: NodeId must be a positive number
type Edge struct {
    NodeId int
    Cost int
}
