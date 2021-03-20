package stringsearch_test

import (
	"app/main/stringsearch"
	"testing"
)

func BenchmarkStringBruteforce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringsearch.SearchBruteForce("HelloHelleHelli", "Helli")
	}
}

func BenchmarkStringKMP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringsearch.SearchKMP("HelloHelleHelli", "Helli")
	}
}

func TestStringsearch(t *testing.T) {
	t.Run("Check if bruteforce works as expected", func(t *testing.T) {
		got := stringsearch.SearchBruteForce("HelloHelleHelli", "Helli")
		expect := 10
		if expect != got {
			t.Errorf("expect %v but got %v", expect, got)
		}
	})

	// t.Run("Check if bruteforce fails as expected", func(t *testing.T) {
	// 	got := stringsearch.SearchBruteForce("Hello", "le")
	// 	expect := -1
	// 	if expect != got {
	// 		t.Errorf("expect %v but got %v", expect, got)
	// 	}
	// })

	t.Run("Check if KMP works as expected", func(t *testing.T) {
		got := stringsearch.SearchKMP("HelloHelleHelli", "Helli")
		expect := 10
		if expect != got {
			t.Errorf("expect %v but got %v", expect, got)
		}
	})

	// t.Run("Check if KMP fails as expected", func(t *testing.T) {
	// 	got := stringsearch.SearchKMP("Hello", "le")
	// 	expect := -1
	// 	if expect != got {
	// 		t.Errorf("expect %v but got %v", expect, got)
	// 	}
	// })
}
