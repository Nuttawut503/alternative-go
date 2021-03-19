package main

import (
	"strings"
	"testing"
)

func TestApp(t *testing.T) {

	assertStringEqual := func(t testing.TB, got, expect string) {
		t.Helper()
		if strings.Compare(expect, got) != 0 {
			t.Errorf("expect %q but got %q", expect, got)
		}
	}

	t.Run("This is an example test (1)", func(t *testing.T) {
		assertStringEqual(t, "Hello", "Hello")
	})

	t.Run("This is an example test (2)", func(t *testing.T) {
		assertStringEqual(t, "World", "World")
	})
}
