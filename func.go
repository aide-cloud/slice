package slice

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// Length returns the length of the given slice.
//
// Parameters:
//   - s: The slice whose length is to be determined.
//
// Returns:
//
//	The number of elements in the slice.
func Length[T any](s []T) int {
	return len(s)
}

// Map applies a transformation function to each element of a slice and returns a new slice with the transformed elements.
//
// Parameters:
//   - list: The original slice.
//   - f: A function that takes an element of type T and its index, and returns an element of type K.
//
// Returns:
//
//	A new slice containing the results of applying the transformation function to each element.
func Map[T, K any](list []T, f func(T, int) K) []K {
	newData := make([]K, 0, len(list))
	for index, v := range list {
		newData = append(newData, f(v, index))
	}
	return newData
}

// Unique returns a new slice with unique elements based on a key function.
//
// Parameters:
//   - s: The original slice.
//   - f: A function that extracts a key from each element of type T. The key must be of a comparable type.
//
// Returns:
//
//	A new slice containing only the first occurrence of each unique key.
func Unique[T any, K comparable](s []T, f func(T) K) []T {
	m := make(map[K]struct{}, len(s))
	indexList := make([]T, 0, len(s))
	for _, v := range s {
		k := f(v)
		if _, ok := m[k]; ok {
			continue
		}
		m[k] = struct{}{}
		indexList = append(indexList, v)
	}
	return indexList
}

// Concat concatenates multiple slices into one.
//
// Parameters:
//   - slices: A variadic parameter representing multiple slices to concatenate.
//
// Returns:
//
//	A single slice containing all elements from the input slices.
func Concat[T any](slices ...[]T) []T {
	if len(slices) == 0 {
		return nil
	}
	s := slices[0]
	for _, slice := range slices[1:] {
		s = append(s, slice...)
	}
	return s
}

// CopyWithIn creates a new slice containing elements at specified indices from the original slice.
//
// Parameters:
//   - s: The original slice.
//   - indexList: A variadic parameter representing the indices of elements to include in the new slice.
//
// Returns:
//
//	A new slice containing elements at the specified indices.
func CopyWithIn[T any](s []T, indexList ...int) []T {
	if len(indexList) == 0 {
		return nil
	}
	newData := make([]T, 0, len(indexList))
	for _, index := range indexList {
		newData = append(newData, s[index])
	}
	return newData
}

// Every checks if all elements in the slice satisfy a given predicate function.
//
// Parameters:
//   - s: The slice to check.
//   - f: A predicate function that takes an element and returns a boolean.
//
// Returns:
//
//	true if all elements satisfy the predicate, false otherwise.
func Every[T any](s []T, f func(T) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}

// Find searches for the first element in the slice that satisfies a given predicate function.
//
// Parameters:
//   - s: The slice to search.
//   - f: A predicate function that takes an element and returns a boolean.
//
// Returns:
//
//	The first element that satisfies the predicate, or the zero value if no such element exists.
func Find[T any](s []T, f func(T) bool) (v T) {
	for _, item := range s {
		if f(item) {
			return item
		}
	}
	return
}

// FindIndex finds the index of the first element in the slice that satisfies a given predicate function.
//
// Parameters:
//   - s: The slice to search.
//   - f: A predicate function that takes an element and returns a boolean.
//
// Returns:
//
//	The index of the first element that satisfies the predicate, or -1 if no such element exists.
func FindIndex[T any](s []T, f func(T) bool) int {
	for i, item := range s {
		if f(item) {
			return i
		}
	}
	return -1
}

// FindLast searches for the last element in the slice that satisfies a given predicate function.
//
// Parameters:
//   - s: The slice to search.
//   - f: A predicate function that takes an element and returns a boolean.
//
// Returns:
//
//	The last element that satisfies the predicate, or the zero value if no such element exists.
func FindLast[T any](s []T, f func(T) bool) (v T) {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return s[i]
		}
	}
	return
}

// FindLastIndex finds the index of the last element in the slice that satisfies a given predicate function.
//
// Parameters:
//   - s: The slice to search.
//   - f: A predicate function that takes an element and returns a boolean.
//
// Returns:
//
//	The index of the last element that satisfies the predicate, or -1 if no such element exists.
func FindLastIndex[T any](s []T, f func(T) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return i
		}
	}
	return -1
}

// ForEach iterates over each element in the slice and applies a provided function to it.
//
// Parameters:
//   - s: The slice to iterate over.
//   - f: A function that takes an element and its index as arguments.
func ForEach[T any](s []T, f func(T, int)) {
	list := make([]T, 0, len(s))
	copy(list, s)
	for index, item := range list {
		f(item, index)
	}
}

// Join converts all elements of the slice to strings and joins them with a specified separator.
//
// Parameters:
//   - s: The slice to join.
//   - seps: An optional separator string.
//
// Returns:
//
//	A string formed by joining the string representations of the slice elements.
func Join[T any](s []T, seps ...string) string {
	ss := make([]string, 0, len(s))
	for _, item := range s {
		ss = append(ss, fmt.Sprintf("%v", item))
	}
	sep := ""
	if len(seps) > 0 {
		sep = seps[0]
	}
	return strings.Join(ss, sep)
}

// Slice returns a subset of the slice, starting at the specified begin index and optionally ending at the end index.
//
// Parameters:
//   - s: The original slice.
//   - indexes: A variadic parameter specifying the begin and optionally end and step indices.
//
// Returns:
//
//	A new slice containing the specified subset.
func Slice[T any](s []T, indexes ...int) []T {
	switch len(indexes) {
	case 1:
		begin := indexes[0]
		return s[begin:]
	case 2:
		begin, end := indexes[0], indexes[1]
		return s[begin:end]
	case 3:
		begin, end, step := indexes[0], indexes[1], indexes[2]
		list := make([]T, 0, (end-begin)/step)
		for i := begin; i < end; i += step {
			list = append(list, s[i])
		}
		return list
	default:
		return s
	}
}

// Fill sets all elements of the slice to a specified value, optionally at specified indices.
//
// Parameters:
//   - s: The slice to fill.
//   - value: The value to set.
//   - indexes: An optional variadic parameter specifying the indices to fill.
//
// Returns:
//
//	The modified slice.
func Fill[T any](s []T, value T, indexes ...int) []T {
	if len(indexes) == 0 {
		for i := range s {
			s[i] = value
		}
		return s
	}

	for _, index := range indexes {
		if index >= 0 && index < len(s) {
			s[index] = value
		}
	}
	return s
}

// String converts the slice to a JSON string representation.
//
// Parameters:
//   - s: The slice to convert.
//
// Returns:
//
//	A JSON string representation of the slice, or "[]" if conversion fails.
func String[T any](s []T) string {
	bs, err := json.Marshal(s)
	if err != nil {
		return "[]"
	}
	return string(bs)
}

// At returns the element at the specified index in the slice.
//
// Parameters:
//   - s: The slice.
//   - index: The index of the element to retrieve.
//
// Returns:
//
//	The element at the specified index.
func At[T any](s []T, index int) T {
	return s[index]
}

// Sort sorts the slice based on a comparison function.
//
// Parameters:
//   - s: The slice to sort.
//   - f: A comparison function that determines the order of elements.
//
// Returns:
//
//	The sorted slice.
func Sort[T any](s []T, f func(a, b T) bool) []T {
	sort.Slice(s, func(i, j int) bool {
		return f(s[i], s[j])
	})
	return s
}

// Filter creates a new slice by filtering elements based on a predicate function.
//
// Parameters:
//   - s: The slice to filter.
//   - f: A predicate function that determines whether an element should be included in the new slice.
//     in the new slice.
//
// Returns:
//
//	A new slice containing only the elements that satisfy the predicate.
func Filter[T any](s []T, f func(T, int) bool) []T {
	list := make([]T, 0, len(s))
	for i, item := range s {
		if f(item, i) {
			list = append(list, item)
		}
	}
	return list
}
