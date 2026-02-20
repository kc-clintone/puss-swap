package helpers

import (
	"reflect"
	"testing"
)

//Stack + Basic Utilities

func TestStackBasics(t *testing.T) {
	s := NewStack([]int{1, 2, 3})
	if s.Len() != 3 {
		t.Fatalf("expected len 3, got %d", s.Len())
	}
	if s.IsEmpty() {
		t.Fatalf("stack should not be empty")
	}

	empty := NewStack([]int{})
	if !empty.IsEmpty() {
		t.Fatalf("stack should be empty")
	}
}

func TestIsSorted(t *testing.T) {
	a := NewStack([]int{1, 2, 3})
	if !IsSorted(a) {
		t.Fatalf("expected sorted")
	}

	b := NewStack([]int{3, 2, 1})
	if IsSorted(b) {
		t.Fatalf("expected not sorted")
	}
}

func TestNormalize(t *testing.T) {
	input := []int{40, 10, 30}
	expected := []int{2, 0, 1}
	result := Normalize(input)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

// Operations

func TestSa(t *testing.T) {
	a := NewStack([]int{2, 1})
	Sa(a)
	expected := []int{1, 2}
	if !reflect.DeepEqual(a.Data, expected) {
		t.Fatalf("expected %v, got %v", expected, a.Data)
	}
}

func TestPaPb(t *testing.T) {
	a := NewStack([]int{1, 2})
	b := NewStack([]int{})

	Pb(a, b)
	if !reflect.DeepEqual(a.Data, []int{2}) {
		t.Fatalf("unexpected A after pb: %v", a.Data)
	}
	if !reflect.DeepEqual(b.Data, []int{1}) {
		t.Fatalf("unexpected B after pb: %v", b.Data)
	}

	Pa(a, b)
	if !reflect.DeepEqual(a.Data, []int{1, 2}) {
		t.Fatalf("unexpected A after pa: %v", a.Data)
	}
}

func TestRotate(t *testing.T) {
	a := NewStack([]int{1, 2, 3})
	Ra(a)
	if !reflect.DeepEqual(a.Data, []int{2, 3, 1}) {
		t.Fatalf("ra failed: %v", a.Data)
	}

	Rra(a)
	if !reflect.DeepEqual(a.Data, []int{1, 2, 3}) {
		t.Fatalf("rra failed: %v", a.Data)
	}
}

// Small Algorithms

func TestSortThree(t *testing.T) {
	cases := [][]int{
		{3, 2, 1},
		{2, 3, 1},
		{3, 1, 2},
		{2, 1, 3},
	}

	for _, c := range cases {
		a := NewStack(c)
		sortThree(a)

		if !IsSorted(a) {
			t.Fatalf("sortThree failed for %v", c)
		}
	}
}

func TestFindMinIndex(t *testing.T) {
	a := NewStack([]int{5, 3, 9, 1})
	idx := findMinIndex(a)
	if idx != 3 {
		t.Fatalf("expected 3, got %d", idx)
	}
}

func TestFindMaxIndex(t *testing.T) {
	a := NewStack([]int{1, 8, 3, 4})
	idx := findMaxIndex(a)
	if idx != 1 {
		t.Fatalf("expected 1, got %d", idx)
	}
}

func TestSortUpToSix(t *testing.T) {
	a := NewStack([]int{4, 1, 6, 2, 5})
	b := NewStack([]int{})

	sortUpToSix(a, b)

	if !IsSorted(a) {
		t.Fatalf("sortUpToSix did not sort")
	}
	if !b.IsEmpty() {
		t.Fatalf("stack B should be empty")
	}
}

// Radix / Chunk Sorting

func TestRadixSort(t *testing.T) {
	input := []int{10, 3, 7, 1, 9, 2, 6, 4, 8, 5}
	a := NewStack(input)
	b := NewStack([]int{})

	radixSort(a, b)

	if !IsSorted(a) {
		t.Fatalf("radixSort failed")
	}
	if !b.IsEmpty() {
		t.Fatalf("stack B should be empty")
	}
}

// Sort Entry Point
func TestSortDispatcher(t *testing.T) {
	cases := [][]int{
		{},
		{2, 1},
		{3, 2, 1},
		{5, 2, 4, 1},
		{9, 3, 7, 1, 8, 2, 6, 4, 5},
	}

	for _, c := range cases {
		a := NewStack(c)
		b := NewStack([]int{})

		Sort(a, b)

		if !IsSorted(a) {
			t.Fatalf("Sort failed for %v", c)
		}
	}
}

// Parser

func TestParseArgsBasic(t *testing.T) {
	nums, err := ParseArgs([]string{"1 2 3"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(nums, expected) {
		t.Fatalf("expected %v, got %v", expected, nums)
	}
}

func TestParseArgsDuplicate(t *testing.T) {
	_, err := ParseArgs([]string{"1 2 2"})
	if err == nil {
		t.Fatalf("expected duplicate error")
	}
}

func TestParseArgsInvalid(t *testing.T) {
	_, err := ParseArgs([]string{"1 a 3"})
	if err == nil {
		t.Fatalf("expected invalid integer error")
	}
}

func TestParseArgsPlaceholderDeterministic(t *testing.T) {
	a, err1 := ParseArgs([]string{"<5 random numbers>"})
	b, err2 := ParseArgs([]string{"<5 random numbers>"})

	if err1 != nil || err2 != nil {
		t.Fatalf("unexpected error")
	}

	if !reflect.DeepEqual(a, b) {
		t.Fatalf("placeholder should be deterministic")
	}
}