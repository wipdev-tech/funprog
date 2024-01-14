package funprog_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/wipdev-tech/funprog"
)

func TestMap(t *testing.T) {
	s1 := []int{1, 2, 3, 5}
	f1 := func(x int) int {
		return x + 1
	}
	e1 := []int{2, 3, 4, 6}
	r1 := funprog.Map(f1, s1)

	if !slices.Equal(r1, e1) {
		t.Fatalf("Expected %v, got %v", e1, r1)
	}

	s2 := []string{"hello", "world"}
	e2 := []string{"HELLO", "WORLD"}
	r2 := funprog.Map(strings.ToUpper, s2)

	if !slices.Equal(r2, e2) {
		t.Fatalf("Expected %v, got %v", e2, r2)
	}
}

func TestFilter(t *testing.T) {
	s1 := []int32{1, 2, 3, 5}
	f1 := func(x int32) bool {
		return x%2 == 0
	}
	e1 := []int32{2}
	r1 := funprog.Filter(f1, s1)

	if !slices.Equal(r1, e1) {
		t.Fatalf("Expected %v, got %v", e1, r1)
	}

	s2 := []string{"hello", "world", "hi"}
	f2 := func(s string) bool {
		return strings.HasPrefix(s, "h")
	}
	e2 := []string{"hello", "hi"}
	r2 := funprog.Filter(f2, s2)

	if !slices.Equal(r2, e2) {
		t.Fatalf("Expected %v, got %v", e2, r2)
	}
}
