package dijkstra

import "github.com/ichinaski/dijkstra/pqueue"

/**
 * NodeId represents a position within the Graph.
 * Note: This type must be comparable - http://golang.org/ref/spec#Comparison_operators
 */
type NodeId interface{}

type Graph interface {
    GetChildren(nodeId NodeId) map[NodeId]int
}

type State struct {
    nodeId NodeId
    parentId NodeId
}

/**
 * Disjktra implementation
 */
func Dijkstra(graph Graph, start, goal NodeId) []NodeId {
    openList := &pqueue.PQueue{}// Nodes not visited yet
    closedList := make(map[NodeId]NodeId)// Visited nodes, and their parents

    state := &State { start, nil }
    openList.Push(state, 0)

    found := false
    for openList.Len() > 0 {
        item, cost := openList.Pop()
        state = item.(*State)

        // Only consider non expanded nodes (not present in closedList)
        if _, ok := closedList[state.nodeId]; !ok {
            // Store this node in the closed list, with a reference to its parent
            closedList[state.nodeId] = state.parentId

            if state.nodeId == goal {
                found = true
                break
            }

            // Add the nodes not present in the closedList into the openList
            children := graph.GetChildren(state.nodeId)
            for childId, childCost := range children {
                if _, ok := closedList[childId]; !ok {
                    childState := &State { childId, state.nodeId }
                    openList.Push(childState, cost + childCost)
                }
            }
        }
    }

    // Build the path, fetching all the nodes in a descendant way, from goal to start
    path := make([]NodeId, 0, len(closedList))
    if found {
        nodeId := goal
        for nodeId != nil {
            path = append(path, nodeId)
            nodeId = closedList[nodeId]
        }

        // Reverse the slice
        for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
            path[i], path[j] = path[j], path[i]
        }
    }

    return path
}
