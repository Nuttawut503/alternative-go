package graph

func BreadthFirstSearch(graph [][]int, startNode int) []int {
	travelOrder := make([]int, 0, len(graph))
	visitted := make([]bool, len(graph))
	nextVisit := make([]int, 0, len(graph))
	nextVisit = append(nextVisit, startNode)
	for len(nextVisit) != 0 {
		visittingNode := nextVisit[0]
		visitted[visittingNode] = true
		travelOrder = append(travelOrder, visittingNode)

		nextVisit = nextVisit[1:]

		for adjacencyNode, isConnected := range graph[visittingNode] {
			if isConnected == 1 && !visitted[adjacencyNode] {
				nextVisit = append(nextVisit, adjacencyNode)
			}
		}
	}
	return travelOrder
}
