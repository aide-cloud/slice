package slice_test

import (
	"reflect"
	"testing"

	"github.com/aide-cloud/slice"
)

// TestString 方法包含多个子测试，用于验证 advancedSlice.String 方法的不同行为
func TestAdvancedSliceString(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		expected string
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			expected: "[]",
		},
		{
			name:     "NonEmptySlice",
			data:     []interface{}{1, 2, 3},
			expected: "[1,2,3]",
		},
		{
			name:     "MixedDataTypes",
			data:     []interface{}{1, "two", 3.0},
			expected: "[1,\"two\",3]",
		},
		{
			name:     "UnconvertibleData",
			data:     []interface{}{map[interface{}]interface{}{1: "one"}},
			expected: "[]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			actual := s.String()
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestMap 方法包含多个子测试，用于验证 advancedSlice.Map 方法的不同行为
func TestAdvancedSliceMap(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		f        func(interface{}, int) interface{}
		expected []interface{}
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			f:        func(x interface{}, i int) interface{} { return x },
			expected: []interface{}{},
		},
		{
			name:     "NonEmptySlice",
			data:     []interface{}{1, 2, 3},
			f:        func(x interface{}, i int) interface{} { return x.(int) + 1 },
			expected: []interface{}{2, 3, 4},
		},
		{
			name:     "MixedDataTypes",
			data:     []interface{}{1, "two", 3.0},
			f:        func(x interface{}, i int) interface{} { return x },
			expected: []interface{}{1, "two", 3.0},
		},
		{
			name:     "ModifyElements",
			data:     []interface{}{1, 2, 3},
			f:        func(x interface{}, i int) interface{} { return x.(int) * 2 },
			expected: []interface{}{2, 4, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			result := s.Map(tt.f)
			actual := result.Values()
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSliceUnique 方法包含多个子测试，用于验证 advancedSlice.Unique 方法的不同行为
func TestAdvancedSliceUnique(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		f        func(interface{}) string
		expected []interface{}
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			f:        func(x interface{}) string { return "" },
			expected: []interface{}{},
		},
		{
			name:     "NonEmptySlice",
			data:     []interface{}{"apple", "banana", "cherry"},
			f:        func(x interface{}) string { return x.(string) },
			expected: []interface{}{"apple", "banana", "cherry"},
		},
		{
			name:     "AllElementsSameKey",
			data:     []interface{}{"apple", "apple", "apple"},
			f:        func(x interface{}) string { return "same" },
			expected: []interface{}{"apple"},
		},
		{
			name:     "DifferentKeys",
			data:     []interface{}{"apple", "banana", "cherry"},
			f:        func(x interface{}) string { return x.(string) },
			expected: []interface{}{"apple", "banana", "cherry"},
		},
		{
			name:     "MixedKeys",
			data:     []interface{}{"apple", "banana", "apple", "cherry", "banana"},
			f:        func(x interface{}) string { return x.(string) },
			expected: []interface{}{"apple", "banana", "cherry"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			result := s.Unique(tt.f)
			actual := result.Values()
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSliceConcat 方法包含多个子测试，用于验证 advancedSlice.Concat 方法的不同行为
func TestAdvancedSliceConcat(t *testing.T) {
	tests := []struct {
		name     string
		initial  []interface{}
		slices   []slice.IAdvancedSlice[interface{}]
		expected []interface{}
	}{
		{
			name:     "EmptySlices",
			initial:  []interface{}{},
			slices:   []slice.IAdvancedSlice[interface{}]{},
			expected: []interface{}{},
		},
		{
			name:     "SingleSlice",
			initial:  []interface{}{1, 2, 3},
			slices:   []slice.IAdvancedSlice[interface{}]{},
			expected: []interface{}{1, 2, 3},
		},
		{
			name:    "MultipleSlices",
			initial: []interface{}{1, 2, 3},
			slices: []slice.IAdvancedSlice[interface{}]{
				slice.NewAdvancedSlice[interface{}](4, 5, 6),
				slice.NewAdvancedSlice[interface{}](7, 8, 9),
			},
			expected: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:    "MixedEmptyAndNonEmptySlices",
			initial: []interface{}{},
			slices: []slice.IAdvancedSlice[interface{}]{
				slice.NewAdvancedSlice[interface{}](1, 2, 3),
				slice.NewAdvancedSlice[interface{}](4, 5),
			},
			expected: []interface{}{1, 2, 3, 4, 5},
		},
		{
			name:    "AllEmptySlices",
			initial: []interface{}{},
			slices: []slice.IAdvancedSlice[interface{}]{
				slice.NewAdvancedSlice[interface{}](4, 5, 6),
				slice.NewAdvancedSlice[interface{}](7, 8, 9),
			},
			expected: []interface{}{4, 5, 6, 7, 8, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.initial...)
			result := s.Concat(tt.slices...)
			actual := result.Values()
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSliceCopyWithIn 方法包含多个子测试，用于验证 advancedSlice.CopyWithIn 方法的不同行为
func TestAdvancedSliceCopyWithIn(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		indexes  []int
		expected []interface{}
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			indexes:  []int{},
			expected: nil,
		},
		{
			name:     "NoIndexes",
			data:     []interface{}{1, 2, 3},
			indexes:  []int{},
			expected: nil,
		},
		{
			name:     "SingleIndex",
			data:     []interface{}{1, 2, 3},
			indexes:  []int{1},
			expected: []interface{}{2},
		},
		{
			name:     "MultipleIndexes",
			data:     []interface{}{1, 2, 3, 4, 5},
			indexes:  []int{0, 2, 4},
			expected: []interface{}{1, 3, 5},
		},
		{
			name:     "OutofBoundsIndexes",
			data:     []interface{}{1, 2, 3},
			indexes:  []int{0, 2, 5},
			expected: []interface{}{1, 3},
		},
		{
			name:     "DuplicateIndexes",
			data:     []interface{}{1, 2, 3},
			indexes:  []int{0, 1, 1, 2},
			expected: []interface{}{1, 2, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			result := s.CopyWithIn(tt.indexes...)
			actual := result.Values()
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSliceEvery 方法包含多个子测试，用于验证 advancedSlice.Every 方法的不同行为
func TestAdvancedSliceEvery(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		f        func(interface{}) bool
		expected bool
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			f:        func(x interface{}) bool { return true },
			expected: true,
		},
		{
			name:     "AllTrue",
			data:     []interface{}{1, 2, 3},
			f:        func(x interface{}) bool { return x.(int) > 0 },
			expected: true,
		},
		{
			name:     "SomeFalse",
			data:     []interface{}{1, 2, 3},
			f:        func(x interface{}) bool { return x.(int) < 3 },
			expected: false,
		},
		{
			name:     "AllFalse",
			data:     []interface{}{-1, -2, -3},
			f:        func(x interface{}) bool { return x.(int) > 0 },
			expected: false,
		},
		{
			name:     "MixedTypes",
			data:     []interface{}{1, "two", 3.0},
			f:        func(x interface{}) bool { return true },
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			actual := s.Every(tt.f)
			if actual != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSliceFind 方法包含多个子测试，用于验证 advancedSlice.Find 方法的不同行为
func TestAdvancedSliceFind(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		f        func(interface{}) bool
		expected interface{}
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			f:        func(x interface{}) bool { return true },
			expected: nil,
		},
		{
			name:     "ElementFound",
			data:     []interface{}{1, 2, 3},
			f:        func(x interface{}) bool { return x.(int) == 2 },
			expected: 2,
		},
		{
			name:     "ElementNotFound",
			data:     []interface{}{1, 2, 3},
			f:        func(x interface{}) bool { return x.(int) == 4 },
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			actual := s.Find(tt.f)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSliceFindIndex 方法包含多个子测试，用于验证 advancedSlice.FindIndex 方法的不同行为
func TestAdvancedSliceFindIndex(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		f        func(interface{}) bool
		expected int
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			f:        func(x interface{}) bool { return true },
			expected: -1,
		},
		{
			name:     "ElementFound",
			data:     []interface{}{1, 2, 3},
			f:        func(x interface{}) bool { return x.(int) == 2 },
			expected: 1,
		},
		{
			name:     "ElementNotFound",
			data:     []interface{}{1, 2, 3},
			f:        func(x interface{}) bool { return x.(int) == 4 },
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			actual := s.FindIndex(tt.f)
			if actual != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSliceFindLast 方法包含多个子测试，用于验证 advancedSlice.FindLast 方法的不同行为
func TestAdvancedSliceFindLast(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		f        func(interface{}) bool
		expected interface{}
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			f:        func(x interface{}) bool { return true },
			expected: nil,
		},
		{
			name:     "ElementFound",
			data:     []interface{}{1, 2, 2, 3},
			f:        func(x interface{}) bool { return x.(int) == 2 },
			expected: 2,
		},
		{
			name:     "ElementNotFound",
			data:     []interface{}{1, 2, 3},
			f:        func(x interface{}) bool { return x.(int) == 4 },
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			actual := s.FindLast(tt.f)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSliceFindLastIndex 方法包含多个子测试，用于验证 advancedSlice.FindLastIndex 方法的不同行为
func TestAdvancedSliceFindLastIndex(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		f        func(interface{}) bool
		expected int
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			f:        func(x interface{}) bool { return true },
			expected: -1,
		},
		{
			name:     "ElementFound",
			data:     []interface{}{1, 2, 2, 3},
			f:        func(x interface{}) bool { return x.(int) == 2 },
			expected: 2,
		},
		{
			name:     "ElementNotFound",
			data:     []interface{}{1, 2, 3},
			f:        func(x interface{}) bool { return x.(int) == 4 },
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			actual := s.FindLastIndex(tt.f)
			if actual != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSliceForEach 方法包含多个子测试，用于验证 advancedSlice.ForEach 方法的不同行为
func TestAdvancedSliceForEach(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		f        func(interface{}, int)
		expected []interface{}
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			f:        func(x interface{}, i int) {},
			expected: []interface{}{},
		},
		{
			name:     "NonEmptySlice",
			data:     []interface{}{1, 2, 3},
			f:        func(x interface{}, i int) {},
			expected: []interface{}{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			s.ForEach(tt.f)
			actual := s.Values()
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSliceJoin 方法包含多个子测试，用于验证 advancedSlice.Join 方法的不同行为
func TestAdvancedSliceJoin(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		sep      []string
		expected string
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			sep:      []string{","},
			expected: "",
		},
		{
			name:     "NonEmptySlice",
			data:     []interface{}{1, 2, 3},
			sep:      []string{","},
			expected: "1,2,3",
		},
		{
			name:     "NoSeparator",
			data:     []interface{}{1, 2, 3},
			sep:      []string{},
			expected: "123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			actual := s.Join(tt.sep...)
			if actual != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSliceSlice 方法包含多个子测试，用于验证 advancedSlice.Slice 方法的不同行为
func TestAdvancedSliceSlice(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		index    []int
		expected []interface{}
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			index:    []int{0, 2},
			expected: []interface{}{},
		},
		{
			name:     "FullSlice",
			data:     []interface{}{1, 2, 3},
			index:    []int{},
			expected: []interface{}{1, 2, 3},
		},
		{
			name:     "PartialSlice",
			data:     []interface{}{1, 2, 3, 4, 5},
			index:    []int{1, 4},
			expected: []interface{}{2, 3, 4},
		},
		{
			name:     "SingleElement",
			data:     []interface{}{1, 2, 3},
			index:    []int{1, 2},
			expected: []interface{}{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			result := s.Slice(tt.index...)
			actual := result.Values()
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSliceFill 方法包含多个子测试，用于验证 advancedSlice.Fill 方法的不同行为
func TestAdvancedSliceFill(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		value    interface{}
		index    []int
		expected []interface{}
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			value:    0,
			index:    []int{},
			expected: []interface{}{},
		},
		{
			name:     "FillAll",
			data:     []interface{}{1, 2, 3},
			value:    0,
			index:    []int{},
			expected: []interface{}{0, 0, 0},
		},
		{
			name:     "FillPartial",
			data:     []interface{}{1, 2, 3, 4, 5},
			value:    0,
			index:    []int{1, 4},
			expected: []interface{}{1, 0, 0, 0, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			result := s.Fill(tt.value, tt.index...)
			actual := result.Values()
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSliceAt 方法包含多个子测试，用于验证 advancedSlice.At 方法的不同行为
func TestAdvancedSliceAt(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		index    int
		expected interface{}
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			index:    0,
			expected: nil,
		},
		{
			name:     "ValidIndex",
			data:     []interface{}{1, 2, 3},
			index:    1,
			expected: 2,
		},
		{
			name:     "OutofBoundsIndex",
			data:     []interface{}{1, 2, 3},
			index:    5,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			actual := s.At(tt.index)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSliceSort 方法包含多个子测试，用于验证 advancedSlice.Sort 方法的不同行为
func TestAdvancedSliceSort(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		f        func(interface{}, interface{}) bool
		expected []interface{}
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			f:        func(x, y interface{}) bool { return x.(int) < y.(int) },
			expected: []interface{}{},
		},
		{
			name:     "NonEmptySlice",
			data:     []interface{}{3, 1, 2},
			f:        func(x, y interface{}) bool { return x.(int) < y.(int) },
			expected: []interface{}{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			result := s.Sort(tt.f)
			actual := result.Values()
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSliceValues 方法包含多个子测试，用于验证 advancedSlice.Values 方法的不同行为
func TestAdvancedSliceValues(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		expected []interface{}
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			expected: []interface{}{},
		},
		{
			name:     "NonEmptySlice",
			data:     []interface{}{1, 2, 3},
			expected: []interface{}{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			actual := s.Values()
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSliceLength 方法包含多个子测试，用于验证 advancedSlice.Length 方法的不同行为
func TestAdvancedSliceLength(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		expected int
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			expected: 0,
		},
		{
			name:     "NonEmptySlice",
			data:     []interface{}{1, 2, 3},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			actual := s.Length()
			if actual != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSliceFilter 方法包含多个子测试，用于验证 advancedSlice.Filter 方法的不同行为
func TestAdvancedSliceFilter(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		f        func(interface{}, int) bool
		expected []interface{}
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			f:        func(x interface{}, i int) bool { return true },
			expected: []interface{}{},
		},
		{
			name:     "NonEmptySlice",
			data:     []interface{}{1, 2, 3, 4, 5},
			f:        func(x interface{}, i int) bool { return x.(int)%2 == 0 },
			expected: []interface{}{2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			actual := s.Filter(tt.f)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSlicePop 方法包含多个子测试，用于验证 advancedSlice.Pop 方法的不同行为
func TestAdvancedSlicePop(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		expected interface{}
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			expected: nil,
		},
		{
			name:     "NonEmptySlice",
			data:     []interface{}{1, 2, 3},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			actual := s.Pop()
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSlicePopIs 方法包含多个子测试，用于验证 advancedSlice.PopIs 方法的不同行为
func TestAdvancedSlicePopIs(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		expected interface{}
		success  bool
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			expected: nil,
			success:  false,
		},
		{
			name:     "NonEmptySlice",
			data:     []interface{}{1, 2, 3},
			expected: 3,
			success:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			actual, success := s.PopIs()
			if success != tt.success {
				t.Errorf("Expected success %v, got %v", tt.success, success)
			}
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSlicePush 方法包含多个子测试，用于验证 advancedSlice.Push 方法的不同行为
func TestAdvancedSlicePush(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		values   []interface{}
		expected []interface{}
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			values:   []interface{}{1, 2, 3},
			expected: []interface{}{1, 2, 3},
		},
		{
			name:     "NonEmptySlice",
			data:     []interface{}{1, 2},
			values:   []interface{}{3, 4},
			expected: []interface{}{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			result := s.Push(tt.values...)
			actual := result.Values()
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSlicePushSlice 方法包含多个子测试，用于验证 advancedSlice.PushSlice 方法的不同行为
func TestAdvancedSlicePushSlice(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		slices   []slice.IAdvancedSlice[interface{}]
		expected []interface{}
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			slices:   []slice.IAdvancedSlice[interface{}]{slice.NewAdvancedSlice[interface{}](1, 2, 3)},
			expected: []interface{}{1, 2, 3},
		},
		{
			name:     "NonEmptySlice",
			data:     []interface{}{1, 2},
			slices:   []slice.IAdvancedSlice[interface{}]{slice.NewAdvancedSlice[interface{}](3, 4), slice.NewAdvancedSlice[interface{}](5, 6)},
			expected: []interface{}{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			result := s.PushSlice(tt.slices...)
			actual := result.Values()
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSliceShift 方法包含多个子测试，用于验证 advancedSlice.Shift 方法的不同行为
func TestAdvancedSliceShift(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		expected interface{}
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			expected: nil,
		},
		{
			name:     "NonEmptySlice",
			data:     []interface{}{1, 2, 3},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			actual := s.Shift()
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSliceShiftIs 方法包含多个子测试，用于验证 advancedSlice.ShiftIs 方法的不同行为
func TestAdvancedSliceShiftIs(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		expected interface{}
		success  bool
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			expected: nil,
			success:  false,
		},
		{
			name:     "NonEmptySlice",
			data:     []interface{}{1, 2, 3},
			expected: 1,
			success:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			actual, success := s.ShiftIs()
			if success != tt.success {
				t.Errorf("Expected success %v, got %v", tt.success, success)
			}
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSliceUnshift 方法包含多个子测试，用于验证 advancedSlice.Unshift 方法的不同行为
func TestAdvancedSliceUnshift(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		values   []interface{}
		expected []interface{}
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			values:   []interface{}{1, 2, 3},
			expected: []interface{}{1, 2, 3},
		},
		{
			name:     "NonEmptySlice",
			data:     []interface{}{4, 5},
			values:   []interface{}{1, 2, 3},
			expected: []interface{}{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			result := s.Unshift(tt.values...)
			actual := result.Values()
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSliceUnshiftSlice 方法包含多个子测试，用于验证 advancedSlice.UnshiftSlice 方法的不同行为
func TestAdvancedSliceUnshiftSlice(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		slices   []slice.IAdvancedSlice[interface{}]
		expected []interface{}
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			slices:   []slice.IAdvancedSlice[interface{}]{slice.NewAdvancedSlice[interface{}](1, 2, 3)},
			expected: []interface{}{1, 2, 3},
		},
		{
			name:     "NonEmptySlice",
			data:     []interface{}{4, 5},
			slices:   []slice.IAdvancedSlice[interface{}]{slice.NewAdvancedSlice[interface{}](1, 2), slice.NewAdvancedSlice[interface{}](3)},
			expected: []interface{}{3, 1, 2, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			result := s.UnshiftSlice(tt.slices...)
			actual := result.Values()
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSliceReverse 方法包含多个子测试，用于验证 advancedSlice.Reverse 方法的不同行为
func TestAdvancedSliceReverse(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		expected []interface{}
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			expected: []interface{}{},
		},
		{
			name:     "NonEmptySlice",
			data:     []interface{}{1, 2, 3},
			expected: []interface{}{3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			result := s.Reverse()
			actual := result.Values()
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSliceRemove 方法包含多个子测试，用于验证 advancedSlice.Remove 方法的不同行为
func TestAdvancedSliceRemove(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		f        func(interface{}, int) bool
		expected []interface{}
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			f:        func(x interface{}, i int) bool { return true },
			expected: []interface{}{},
		},
		{
			name:     "RemoveAll",
			data:     []interface{}{1, 2, 3},
			f:        func(x interface{}, i int) bool { return true },
			expected: []interface{}{},
		},
		{
			name:     "RemoveEvenIndices",
			data:     []interface{}{1, 2, 3, 4, 5},
			f:        func(x interface{}, i int) bool { return i%2 == 0 },
			expected: []interface{}{2, 4},
		},
		{
			name:     "RemoveOddValues",
			data:     []interface{}{1, 2, 3, 4, 5},
			f:        func(x interface{}, i int) bool { return x.(int)%2 != 0 },
			expected: []interface{}{2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			result := s.Remove(tt.f)
			actual := result.Values()
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// TestAdvancedSliceRemoveAt 方法包含多个子测试，用于验证 advancedSlice.RemoveAt 方法的不同行为
func TestAdvancedSliceRemoveAt(t *testing.T) {
	tests := []struct {
		name     string
		data     []interface{}
		index    int
		expected []interface{}
	}{
		{
			name:     "EmptySlice",
			data:     []interface{}{},
			index:    0,
			expected: []interface{}{},
		},
		{
			name:     "RemoveFirstElement",
			data:     []interface{}{1, 2, 3},
			index:    0,
			expected: []interface{}{2, 3},
		},
		{
			name:     "RemoveMiddleElement",
			data:     []interface{}{1, 2, 3, 4},
			index:    2,
			expected: []interface{}{1, 2, 4},
		},
		{
			name:     "RemoveLastElement",
			data:     []interface{}{1, 2, 3},
			index:    2,
			expected: []interface{}{1, 2},
		},
		{
			name:     "IndexOutOfBounds",
			data:     []interface{}{1, 2, 3},
			index:    5,
			expected: []interface{}{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.NewAdvancedSlice(tt.data...)
			result := s.RemoveAt(tt.index)
			actual := result.Values()
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
