package dynamic_test

import (
	"app/main/dynamic"
	"reflect"
	"testing"
)

func TestDynamicProgramming(t *testing.T) {
	t.Run("Test fibonacci", func(t *testing.T) {
		got := dynamic.Fibonacci(10)
		expect := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expect %v but got %v", expect, got)
		}
	})

	t.Run("Test LIS", func(t *testing.T) {
		got := dynamic.LongestIncreasingSubsequence([]int{4, 5, 0, 6, 1, 2, 3})
		expect := 4

		if expect != got {
			t.Errorf("expect %v but got %v", expect, got)
		}
	})
}
