package graph_test

import (
	"app/main/algorithms/graph"
	"reflect"
	"testing"
)

func TestGraph(t *testing.T) {
	t.Run("Check if BFS runs correctly", func(t *testing.T) {
		g := [][]int{
			{1, 1, 0, 1, 0},
			{1, 1, 0, 0, 1},
			{0, 0, 1, 1, 0},
			{1, 0, 1, 1, 0},
			{0, 1, 0, 0, 1},
		}
		got := graph.BreadthFirstSearch(g, 0)
		expect := []int{0, 1, 3, 4, 2}

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expect %v but got %v", expect, got)
		}
	})

	t.Run("Check if shortest path gets the correct answer", func(t *testing.T) {
		g := [][]int{
			{0, 10001, 10001, 10001, 9},
			{10001, 0, 3, 2, 10001},
			{10001, 3, 0, 10001, 9},
			{10001, 2, 10001, 0, 7},
			{9, 10001, 9, 7, 0},
		}
		got := graph.ShortestPath(g, 0)
		expect := []int{0, 18, 18, 16, 9}

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expect %v but got %v", expect, got)
		}
	})
}
