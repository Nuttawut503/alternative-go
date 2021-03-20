package sorting_test

import (
	"app/main/sorting"
	"reflect"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	t.Run("Check heapsort", func(t *testing.T) {
		arr := []int{1, 5, 6, 2, 4, 3, 3, 0, 0, 7}
		got := sorting.Heapsort(arr)
		sort.Ints(arr)
		expect := arr

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expect %v but got %v", expect, got)
		}
	})

	t.Run("Check mergesort", func(t *testing.T) {
		arr := []int{1, 5, 6, 2, 4, 3, 3, 0, 0, 7}
		got := sorting.Mergesort(arr)
		sort.Ints(arr)
		expect := arr

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expect %v but got %v", expect, got)
		}
	})

	t.Run("Check quicksort", func(t *testing.T) {
		arr := []int{1, 5, 6, 2, 4, 3, 3, 0, 0, 7}
		got := sorting.Quicksort(arr)
		sort.Ints(arr)
		expect := arr

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expect %v but got %v", expect, got)
		}
	})
}
