package slice_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aide-cloud/slice"
)

func TestLength(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"non-empty slice", []int{1, 2, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slice.Length(tt.s); got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		list     []int
		f        func(int, int) string
		wantList []string
	}{
		{"empty slice", []int{}, func(i int, idx int) string { return fmt.Sprintf("%d", i) }, []string{}},
		{"non-empty slice", []int{1, 2, 3}, func(i int, idx int) string { return fmt.Sprintf("%d", i*idx) }, []string{"0", "2", "6"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotList := slice.Map(tt.list, tt.f); !reflect.DeepEqual(gotList, tt.wantList) {
				t.Errorf("Map() = %v, want %v", gotList, tt.wantList)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		f    func(int) int
		want []int
	}{
		{"empty slice", []int{}, func(i int) int { return i }, []int{}},
		{"non-empty slice", []int{1, 2, 2, 3, 4, 4, 5}, func(i int) int { return i }, []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slice.Unique(tt.s, tt.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcat(t *testing.T) {
	tests := []struct {
		name string
		args [][]int
		want []int
	}{
		{"no slices", nil, nil},
		{"one slice", [][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{"multiple slices", [][]int{{1, 2}, {3, 4}, {5, 6}}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slice.Concat(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Concat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopyWithIn(t *testing.T) {
	tests := []struct {
		name      string
		s         []int
		indexList []int
		want      []int
	}{
		{"empty slice", []int{}, []int{}, nil},
		{"valid indices", []int{1, 2, 3, 4, 5}, []int{0, 2, 4}, []int{1, 3, 5}},
		{"invalid indices", []int{1, 2, 3}, []int{0, 5}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slice.CopyWithIn(tt.s, tt.indexList...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CopyWithIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvery(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		f    func(int) bool
		want bool
	}{
		{"empty slice", []int{}, func(i int) bool { return i > 0 }, true},
		{"all elements satisfy predicate", []int{1, 2, 3}, func(i int) bool { return i > 0 }, true},
		{"not all elements satisfy predicate", []int{1, -1, 3}, func(i int) bool { return i > 0 }, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slice.Every(tt.s, tt.f); got != tt.want {
				t.Errorf("Every() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		f    func(int) bool
		want int
	}{
		{"empty slice", []int{}, func(i int) bool { return i == 2 }, 0},
		{"element found", []int{1, 2, 3}, func(i int) bool { return i == 2 }, 2},
		{"element not found", []int{1, 3, 5}, func(i int) bool { return i == 2 }, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slice.Find(tt.s, tt.f); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindIndex(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		f    func(int) bool
		want int
	}{
		{"empty slice", []int{}, func(i int) bool { return i == 2 }, -1},
		{"element found", []int{1, 2, 3}, func(i int) bool { return i == 2 }, 1},
		{"element not found", []int{1, 3, 5}, func(i int) bool { return i == 2 }, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slice.FindIndex(tt.s, tt.f); got != tt.want {
				t.Errorf("FindIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindLast(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		f    func(int) bool
		want int
	}{
		{"empty slice", []int{}, func(i int) bool { return i == 2 }, 0},
		{"element found", []int{1, 2, 3, 2}, func(i int) bool { return i == 2 }, 2},
		{"element not found", []int{1, 3, 5}, func(i int) bool { return i == 2 }, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slice.FindLast(tt.s, tt.f); got != tt.want {
				t.Errorf("FindLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindLastIndex(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		f    func(int) bool
		want int
	}{
		{"empty slice", []int{}, func(i int) bool { return i == 2 }, -1},
		{"element found", []int{1, 2, 3, 2}, func(i int) bool { return i == 2 }, 3},
		{"element not found", []int{1, 3, 5}, func(i int) bool { return i == 2 }, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slice.FindLastIndex(tt.s, tt.f); got != tt.want {
				t.Errorf("FindLastIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForEach(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		f    func(int, int)
	}{
		{"empty slice", []int{}, func(i int, idx int) {}},
		{"non-empty slice", []int{1, 2, 3}, func(i int, idx int) { t.Logf("Element: %d, Index: %d", i, idx) }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slice.ForEach(tt.s, tt.f)
		})
	}
}

func TestJoin(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		seps []string
		want string
	}{
		{"empty slice", []int{}, nil, ""},
		{"non-empty slice", []int{1, 2, 3}, nil, "123"},
		{"with separator", []int{1, 2, 3}, []string{","}, "1,2,3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slice.Join(tt.s, tt.seps...); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestSlice tests the Slice function with various scenarios.
func TestSlice(t *testing.T) {
	tests := []struct {
		name    string
		s       []int
		indexes []int
		want    []int
	}{
		{"empty slice", []int{}, []int{0}, []int{}},
		{"single index", []int{1, 2, 3, 4, 5}, []int{2}, []int{3, 4, 5}},
		{"two indices", []int{1, 2, 3, 4, 5}, []int{1, 3}, []int{2, 3}},
		{"three indices", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{1, 7, 2}, []int{2, 4, 6, 8}},
		{"begin greater than end", []int{1, 2, 3, 4, 5}, []int{3, 1}, nil},
		{"begin negative", []int{1, 2, 3, 4, 5}, []int{-1, 3}, nil},
		{"end greater than length", []int{1, 2, 3, 4, 5}, []int{1, 10}, []int{2, 3, 4, 5}},
		{"step zero", []int{1, 2, 3, 4, 5}, []int{1, 4, 0}, nil},
		{"step negative", []int{1, 2, 3, 4, 5}, []int{1, 4, -1}, nil},
		{"no indexes", []int{1, 2, 3, 4, 5}, []int{}, []int{1, 2, 3, 4, 5}},
		{"begin equals end", []int{1, 2, 3, 4, 5}, []int{2, 2}, []int{}},
		{"begin equals length", []int{1, 2, 3, 4, 5}, []int{5}, []int{}},
		{"begin greater than length", []int{1, 2, 3, 4, 5}, []int{6}, nil},
		{"begin and end equal length", []int{1, 2, 3, 4, 5}, []int{5, 5}, []int{}},
		{"begin and end greater than length", []int{1, 2, 3, 4, 5}, []int{6, 7}, nil},
		{"step greater than length", []int{1, 2, 3, 4, 5}, []int{0, 5, 6}, []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slice.Slice(tt.s, tt.indexes...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice(%v, %v) = %v, want %v", tt.s, tt.indexes, got, tt.want)
			}
		})
	}
}

func TestFill(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		value    int
		indexes  []int
		expected []int
	}{
		{"fill all", []int{1, 2, 3}, 0, []int{}, []int{0, 0, 0}},
		{"fill from index", []int{1, 2, 3}, 0, []int{1}, []int{1, 0, 0}},
		{"fill range", []int{1, 2, 3, 4, 5}, 0, []int{1, 3}, []int{1, 0, 0, 4, 5}},
		{"fill invalid index", []int{1, 2, 3}, 0, []int{5}, []int{1, 2, 3}},
		{"fill negative index", []int{1, 2, 3}, 0, []int{-1}, []int{3, 0, 0}},
		{"fill reverse range", []int{1, 2, 3, 4, 5}, 0, []int{3, 1}, []int{1, 0, 0, 4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := slice.Fill(tt.slice, tt.value, tt.indexes...)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Fill(%v, %v, %v) = %v, want %v", tt.slice, tt.value, tt.indexes, result, tt.expected)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		expected string
	}{
		{"empty slice", []int{}, "[]"},
		{"non-empty slice", []int{1, 2, 3}, "[1,2,3]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := slice.String(tt.slice)
			if result != tt.expected {
				t.Errorf("String(%v) = %v, want %v", tt.slice, result, tt.expected)
			}
		})
	}
}

func TestAt(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		index    int
		expected int
	}{
		{"valid index", []int{1, 2, 3}, 1, 2},
		{"out of bounds", []int{1, 2, 3}, 5, 0},
		{"negative index", []int{1, 2, 3}, -1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := slice.At(tt.slice, tt.index)
			if result != tt.expected {
				t.Errorf("At(%v, %v) = %v, want %v", tt.slice, tt.index, result, tt.expected)
			}
		})
	}
}

func TestSort(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		less     func(a, b int) bool
		expected []int
	}{
		{"ascending", []int{3, 1, 2}, func(a, b int) bool { return a < b }, []int{1, 2, 3}},
		{"descending", []int{3, 1, 2}, func(a, b int) bool { return a > b }, []int{3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := slice.Sort(tt.slice, tt.less)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Sort(%v, f) = %v, want %v", tt.slice, result, tt.expected)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		filter   func(a int, index int) bool
		expected []int
	}{
		{"filter even", []int{1, 2, 3, 4}, func(a int, index int) bool { return a%2 == 0 }, []int{2, 4}},
		{"filter odd", []int{1, 2, 3, 4}, func(a int, index int) bool { return a%2 != 0 }, []int{1, 3}},
		{"filter all", []int{1, 2, 3, 4}, func(a int, index int) bool { return true }, []int{1, 2, 3, 4}},
		{"filter none", []int{1, 2, 3, 4}, func(a int, index int) bool { return false }, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := slice.Filter(tt.slice, tt.filter)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Filter(%v, f) = %v, want %v", tt.slice, result, tt.expected)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		filter   func(a int, index int) bool
		expected []int
	}{
		{"remove even", []int{1, 2, 3, 4}, func(a int, index int) bool { return a%2 == 0 }, []int{1, 3}},
		{"remove odd", []int{1, 2, 3, 4}, func(a int, index int) bool { return a%2 != 0 }, []int{2, 4}},
		{"remove all", []int{1, 2, 3, 4}, func(a int, index int) bool { return true }, []int{}},
		{"remove none", []int{1, 2, 3, 4}, func(a int, index int) bool { return false }, []int{1, 2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := slice.Remove(tt.slice, tt.filter)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Remove(%v, f) = %v, want %v", tt.slice, result, tt.expected)
			}
		})
	}
}

func TestRemoveAt(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		index    int
		expected []int
	}{
		{"remove middle", []int{1, 2, 3, 4}, 2, []int{1, 2, 4}},
		{"remove first", []int{1, 2, 3, 4}, 0, []int{2, 3, 4}},
		{"remove last", []int{1, 2, 3, 4}, 3, []int{1, 2, 3}},
		{"remove out of bounds", []int{1, 2, 3, 4}, 5, []int{1, 2, 3, 4}},
		{"remove negative index", []int{1, 2, 3, 4}, -1, []int{1, 2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := slice.RemoveAt(tt.slice, tt.index)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("RemoveAt(%v, %v) = %v, want %v", tt.slice, tt.index, result, tt.expected)
			}
		})
	}
}
