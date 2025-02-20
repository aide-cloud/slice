// func_test.go
package slice_test

import (
	"fmt"
	"testing"

	"github.com/aide-cloud/slice"
)

func slicesStringEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func slicesIntEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestLength(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"empty slice", []int{}, 0},
		{"one element", []int{1}, 1},
		{"multiple elements", []int{1, 2, 3, 4, 5}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := slice.Length(tt.input)
			if actual != tt.expected {
				t.Errorf("Length(%v) = %d; want %d", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []string
	}{
		{"empty slice", []int{}, []string{}},
		{"one element", []int{1}, []string{"1"}},
		{"multiple elements", []int{1, 2, 3}, []string{"1", "2", "3"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := slice.Map(tt.input, func(x int, _ int) string { return fmt.Sprintf("%d", x) })
			if !slicesStringEqual(actual, tt.expected) {
				t.Errorf("Map(%v) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{"empty slice", []string{}, []string{}},
		{"no duplicates", []string{"apple", "banana", "cherry"}, []string{"apple", "banana", "cherry"}},
		{"with duplicates", []string{"apple", "banana", "apple", "orange"}, []string{"apple", "banana", "orange"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := slice.Unique(tt.input, func(x string) string { return x })
			if !slicesStringEqual(actual, tt.expected) {
				t.Errorf("Unique(%v) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestConcat(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected []int
	}{
		{"empty slices", [][]int{}, []int{}},
		{"one slice", [][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{"multiple slices", [][]int{{1, 2}, {3, 4}, {5, 6}}, []int{1, 2, 3, 4, 5, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := slice.Concat(tt.input...)
			if !slicesIntEqual(actual, tt.expected) {
				t.Errorf("Concat(%v) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestCopyWithIn(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		indexes  []int
		expected []int
	}{
		{"empty slice", []int{}, []int{}, []int{}},
		{"one index", []int{1, 2, 3}, []int{1}, []int{2}},
		{"multiple indices", []int{1, 2, 3, 4, 5}, []int{0, 2, 4}, []int{1, 3, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := slice.CopyWithIn(tt.input, tt.indexes...)
			if !slicesIntEqual(actual, tt.expected) {
				t.Errorf("CopyWithIn(%v, %v) = %v; want %v", tt.input, tt.indexes, actual, tt.expected)
			}
		})
	}
}

func TestEvery(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		predicate func(int) bool
		expected  bool
	}{
		{"empty slice", []int{}, func(x int) bool { return x > 0 }, true},
		{"all true", []int{1, 2, 3}, func(x int) bool { return x > 0 }, true},
		{"some false", []int{1, 2, -3}, func(x int) bool { return x > 0 }, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := slice.Every(tt.input, tt.predicate)
			if actual != tt.expected {
				t.Errorf(`Every(%v, f) = %v; want %v`, tt.input, actual, tt.expected)
			}
		})
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		predicate func(int) bool
		expected  int
	}{
		{"empty slice", []int{}, func(x int) bool { return x > 0 }, 0},
		{"found", []int{1, 2, 3}, func(x int) bool { return x > 1 }, 2},
		{"not found", []int{1, 2, 3}, func(x int) bool { return x > 3 }, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := slice.Find(tt.input, tt.predicate)
			if actual != tt.expected {
				t.Errorf("Find(%v, f) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestFindIndex(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		predicate func(int) bool
		expected  int
	}{
		{"empty slice", []int{}, func(x int) bool { return x > 0 }, -1},
		{"found", []int{1, 2, 3}, func(x int) bool { return x > 1 }, 1},
		{"not found", []int{1, 2, 3}, func(x int) bool { return x > 3 }, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := slice.FindIndex(tt.input, tt.predicate)
			if actual != tt.expected {
				t.Errorf("FindIndex(%v, f) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestFindLast(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		predicate func(int) bool
		expected  int
	}{
		{"empty slice", []int{}, func(x int) bool { return x > 0 }, 0},
		{"found", []int{1, 2, 3, 2}, func(x int) bool { return x == 2 }, 2},
		{"not found", []int{1, 2, 3}, func(x int) bool { return x > 3 }, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := slice.FindLast(tt.input, tt.predicate)
			if actual != tt.expected {
				t.Errorf("FindLast(%v, f) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestFindLastIndex(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		predicate func(int) bool
		expected  int
	}{
		{"empty slice", []int{}, func(x int) bool { return x > 0 }, -1},
		{"found", []int{1, 2, 3, 2}, func(x int) bool { return x == 2 }, 3},
		{"not found", []int{1, 2, 3}, func(x int) bool { return x > 3 }, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := slice.FindLastIndex(tt.input, tt.predicate)
			if actual != tt.expected {
				t.Errorf("FindLastIndex(%v, f) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestForEach(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty slice", []int{}, []int{}},
		{"modify elements", []int{1, 2, 3}, []int{2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := make([]int, len(tt.input))
			copy(actual, tt.input)
			slice.ForEach(actual, func(x int, i int) {
				if x != actual[i] || x+1 != tt.input[i] {
					t.Errorf("ForEach(%v, f) modified slice incorrectly", tt.input)
				}
			})
		})
	}
}

func TestJoin(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		sep      string
		expected string
	}{
		{"empty slice", []int{}, ",", ""},
		{"one element", []int{1}, ",", "1"},
		{"multiple elements", []int{1, 2, 3}, ",", "1,2,3"},
		{"multiple elements with custom sep", []int{1, 2, 3}, "-", "1-2-3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := slice.Join(tt.input, tt.sep)
			if actual != tt.expected {
				t.Errorf("Join(%v, %v) = %v; want %v", tt.input, tt.sep, actual, tt.expected)
			}
		})
	}
}

func TestSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		indexes  []int
		expected []int
	}{
		{"empty slice", []int{}, []int{0}, []int{}},
		{"one index", []int{1, 2, 3}, []int{1}, []int{2, 3}},
		{"two indices", []int{1, 2, 3, 4, 5}, []int{1, 4}, []int{2, 3, 4}},
		{"three indices", []int{1, 2, 3, 4, 5}, []int{0, 4, 2}, []int{1, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := slice.Slice(tt.input, tt.indexes...)
			if !slicesIntEqual(actual, tt.expected) {
				t.Errorf("Slice(%v, %v) = %v; want %v", tt.input, tt.indexes, actual, tt.expected)
			}
		})
	}
}

func TestFill(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		value    int
		indexes  []int
		expected []int
	}{
		{"empty slice", []int{}, 1, []int{0}, []int{}},
		{"fill all", []int{1, 2, 3}, 0, []int{}, []int{0, 0, 0}},
		{"fill specific indices", []int{1, 2, 3, 4, 5}, 0, []int{1, 3}, []int{1, 0, 3, 0, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := slice.Fill(tt.input, tt.value, tt.indexes...)
			if !slicesIntEqual(actual, tt.expected) {
				t.Errorf("Fill(%v, %v, %v) = %v; want %v", tt.input, tt.value, tt.indexes, actual, tt.expected)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected string
	}{
		{"empty slice", []int{}, "[]"},
		{"one element", []int{1}, "[1]"},
		{"multiple elements", []int{1, 2, 3}, "[1,2,3]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := slice.String(tt.input)
			if actual != tt.expected {
				t.Errorf("String(%v) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}
