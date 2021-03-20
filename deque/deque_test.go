package deque

import (
	"reflect"
	"testing"
)

func TestDeque(t *testing.T) {
	assertIntegerEqual := func(t testing.TB, expect, got int) {
		t.Helper()
		if expect != got {
			t.Errorf("expect %q but got %q", expect, got)
		}
	}

	t.Run("Check if Deque length starts with 0", func(t *testing.T) {
		newDeque := new(Deque)
		got := newDeque.Len()
		expect := 0

		assertIntegerEqual(t, expect, got)
	})

	t.Run("Check if Front-side functions work", func(t *testing.T) {
		newDeque := new(Deque)
		newDeque.PushFront(10)
		newDeque.PopFront()
		newDeque.PushFront(2)
		newDeque.PushFront(1)
		newDeque.PushFront(0)
		newDeque.PopFront()
		dqList := newDeque.ToList()
		got := make([]int, len(dqList))
		for i, v := range dqList {
			got[i] = v.(int)
		}
		expect := []int{1, 2}

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expect %q but got %q", expect, got)
		}
	})

	t.Run("Check if Back-side functions work", func(t *testing.T) {
		newDeque := new(Deque)
		newDeque.PushBack(10)
		newDeque.PopBack()
		newDeque.PushBack(0)
		newDeque.PushBack(1)
		newDeque.PushBack(2)
		newDeque.PopBack()
		dqList := newDeque.ToList()
		got := make([]int, len(dqList))
		for i, v := range dqList {
			got[i] = v.(int)
		}
		expect := []int{0, 1}

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expect %v but got %v", expect, got)
		}
	})
}
