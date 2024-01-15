package funprog_test

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	f "github.com/wipdev-tech/funprog"
)

func hasPrefixH(s string) bool {
	return strings.HasPrefix(s, "h")
}

func isEven(x int) bool {
	return x%2 == 0
}

func TestMap(t *testing.T) {
	s1 := []int{1, 2, 3, 5}
	f1 := func(x int) int {
		return x + 1
	}
	e1 := []int{2, 3, 4, 6}
	r1 := f.Map(f1, s1)

	if !slices.Equal(r1, e1) {
		t.Fatalf("Expected %v, got %v", e1, r1)
	}

	s2 := []string{"hello", "world"}
	e2 := []string{"HELLO", "WORLD"}
	r2 := f.Map(strings.ToUpper, s2)

	if !slices.Equal(r2, e2) {
		t.Fatalf("Expected %v, got %v", e2, r2)
	}
}

func TestFilter(t *testing.T) {
	s1 := []int{1, 2, 3, 5}
	e1 := []int{2}
	r1 := f.Filter(isEven, s1)

	if !slices.Equal(r1, e1) {
		t.Fatalf("Expected %v, got %v", e1, r1)
	}

	s2 := []string{"hello", "world", "hi"}
	e2 := []string{"hello", "hi"}
	r2 := f.Filter(hasPrefixH, s2)

	if !slices.Equal(r2, e2) {
		t.Fatalf("Expected %v, got %v", e2, r2)
	}
}

func TestReduce(t *testing.T) {
	s1 := []int{2, 0, 10, -1, 3}
	f1 := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	e1 := -1
	r1 := f.Reduce(f1, s1)

	if r1 != e1 {
		t.Fatalf("Expected %v, got %v", e1, r1)
	}

	s2 := []string{"hello", "world", ":)"}
	f2 := func(x, y string) string {
		return fmt.Sprintf("%s%c", x, y[0])
	}
	e2 := "hw:"
	r2 := f.Reduce(f2, s2)

	if r2 != e2 {
		t.Fatalf("Expected %v, got %v", e2, r2)
	}
}

func TestComp(t *testing.T) {
	f1 := func(x int) int { return x / 7 }
	f2 := func(x int) int { return x - 1 }
	f3 := func(x int) int { return x / 3 }
	e1 := f3(f2(f1(49)))
	r1 := f.Comp(f3, f2, f1)(49)

	if r1 != e1 {
		t.Fatalf("Expected %v, got %v", e1, r1)
	}

	f4 := func(x string) string {
		return strings.ReplaceAll(x, " ", "_")
	}
	f5 := strings.TrimSpace
	e2 := f5(f4(" hiyaa "))
	r2 := f.Comp(f5, f4)(" hiyaa ")

	if r2 != e2 {
		t.Fatalf("Expected %v, got %v", e2, r2)
	}
}

func TestCompR(t *testing.T) {
	f1 := func(x int) int {
		return x / 7
	}
	f2 := func(x int) int {
		return x - 1
	}
	f3 := func(x int) int {
		return x / 3
	}
	e1 := f3(f2(f1(49)))
	r1 := f.CompR(f1, f2, f3)(49)

	if r1 != e1 {
		t.Fatalf("Expected %v, got %v", e1, r1)
	}

	f4 := func(x string) string {
		return strings.ReplaceAll(x, " ", "_")
	}
	f5 := strings.TrimSpace
	e2 := f5(f4(" hiyaa "))
	r2 := f.CompR(f4, f5)(" hiyaa ")

	if r2 != e2 {
		t.Fatalf("Expected %v, got %v", e2, r2)
	}
}

func TestAny(t *testing.T) {
	s1 := []string{"hello", "lol", "hershey"}
	r1 := f.Any(hasPrefixH, s1)
	if !r1 {
		t.Fatalf("`f.Any(hasPrefixH, s1)` should be true")
	}

	s2 := []int{3, 1, 5}
	r2 := f.Any(isEven, s2)
	if r2 {
		t.Fatalf("`f.Any(isEven, s2)` should be false")
	}
}

func TestAll(t *testing.T) {
	s1 := []string{"hello", "hi", "hershey"}
	r1 := f.All(hasPrefixH, s1)
	if !r1 {
		t.Fatalf("`f.All(hasPrefixH, s1)` should be true")
	}

	s2 := []int{3, 1, 5}
	r2 := f.Any(isEven, s2)
	if r2 {
		t.Fatalf("`f.All(isEven, s2)` should be false")
	}
}

func TestFindIndex(t *testing.T) {
	s1 := []string{"hello", "hi", "hershey"}
	e1 := 0
	r1 := f.FindIndex(hasPrefixH, s1)

	if r1 != e1 {
		t.Fatalf("Expected %v, got %v", e1, r1)
	}

	s2 := []int{1, 3, 5}
	e2 := -1
	r2 := f.FindIndex(isEven, s2)

	if r2 != e2 {
		t.Fatalf("Expected %v, got %v", e2, r2)
	}
}

func TestFindIndices(t *testing.T) {
	s1 := []string{"hello", "hi", "hershey"}
	e1 := []int{0, 1, 2}
	r1 := f.FindIndices(hasPrefixH, s1)

	if !slices.Equal(r1, e1) {
		t.Fatalf("Expected %v, got %v", e1, r1)
	}

	s2 := []int{1, 3, 5}
	e2 := []int{}
	r2 := f.FindIndices(isEven, s2)

	if !slices.Equal(r2, e2) {
		t.Fatalf("Expected %v, got %v", e2, r2)
	}
}
