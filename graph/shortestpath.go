package graph

func getInfinites(size int) []int {
	infs := make([]int, size)
	for i := range infs {
		infs[i] = 10001
	}
	return infs
}

func minIndex(arr []int, blist []bool) int {
	mIndex := -1
	for i := 0; i < len(arr); i++ {
		if !blist[i] && (mIndex == -1 || arr[i] < arr[mIndex]) {
			mIndex = i
		}
	}
	return mIndex
}

func ShortestPath(graph [][]int, from int) []int {
	weightRecord := getInfinites(len(graph))
	visitted := make([]bool, len(graph))
	weightRecord[from] = 0
	for {
		visittingNode := minIndex(weightRecord, visitted)
		if visittingNode == -1 {
			break
		}
		visitted[visittingNode] = true
		for adjacencyNode, weight := range graph[visittingNode] {
			if weightRecord[visittingNode]+weight < weightRecord[adjacencyNode] {
				weightRecord[adjacencyNode] = weightRecord[visittingNode] + weight
			}
		}
	}
	return weightRecord
}
