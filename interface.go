package slice

import (
	"fmt"
)

// IAdvancedSlice defines an interface for advanced slice manipulation.
// It extends the fmt.Stringer interface to provide string representation capabilities.
type IAdvancedSlice[T any] interface {
	fmt.Stringer

	// Length returns the number of elements in the slice.
	//
	// Returns:
	//   The length of the slice as an integer.
	Length() int

	// Map applies a transformation function to each element of the slice and returns a new advanced slice with the transformed elements.
	//
	// Parameters:
	//   - f: A function that takes an element of type T and its index, and returns an element of type T.
	//
	// Returns:
	//   A new IAdvancedSlice[T] containing the results of applying the transformation function to each element.
	Map(f func(T, int) T) IAdvancedSlice[T]

	// Unique returns a new advanced slice with unique elements based on a key function.
	//
	// Parameters:
	//   - f: A function that extracts a key from each element of type T. The key must be a string.
	//
	// Returns:
	//   A new IAdvancedSlice[T] containing only the first occurrence of each unique key.
	Unique(f func(T) string) IAdvancedSlice[T]

	// Concat concatenates multiple slices into one and returns a new advanced slice.
	//
	// Parameters:
	//   - s: A variadic parameter representing multiple slices to concatenate.
	//
	// Returns:
	//   A new IAdvancedSlice[T] containing all elements from the input slices.
	Concat(s ...IAdvancedSlice[T]) IAdvancedSlice[T]

	// CopyWithIn creates a new advanced slice containing elements at specified indices from the original slice.
	//
	// Parameters:
	//   - indexes: A variadic parameter representing the indices of elements to include in the new slice.
	//
	// Returns:
	//   A new IAdvancedSlice[T] containing elements at the specified indices.
	CopyWithIn(indexes ...int) IAdvancedSlice[T]

	// Every checks if all elements in the slice satisfy a given predicate function.
	//
	// Parameters:
	//   - f: A predicate function that takes an element and returns a boolean.
	//
	// Returns:
	//   true if all elements satisfy the predicate, false otherwise.
	Every(f func(T) bool) bool

	// Find searches for the first element in the slice that satisfies a given predicate function.
	//
	// Parameters:
	//   - f: A predicate function that takes an element and returns a boolean.
	//
	// Returns:
	//   The first element that satisfies the predicate, or the zero value if no such element exists.
	Find(f func(T) bool) T

	// FindIndex finds the index of the first element in the slice that satisfies a given predicate function.
	//
	// Parameters:
	//   - f: A predicate function that takes an element and returns a boolean.
	//
	// Returns:
	//   The index of the first element that satisfies the predicate, or -1 if no such element exists.
	FindIndex(f func(T) bool) int

	// FindLast searches for the last element in the slice that satisfies a given predicate function.
	//
	// Parameters:
	//   - f: A predicate function that takes an element and returns a boolean.
	//
	// Returns:
	//   The last element that satisfies the predicate, or the zero value if no such element exists.
	FindLast(f func(T) bool) T

	// FindLastIndex finds the index of the last element in the slice that satisfies a given predicate function.
	//
	// Parameters:
	//   - f: A predicate function that takes an element and returns a boolean.
	//
	// Returns:
	//   The index of the last element that satisfies the predicate, or -1 if no such element exists.
	FindLastIndex(f func(T) bool) int

	// ForEach iterates over each element in the slice and applies a provided function to it.
	//
	// Parameters:
	//   - f: A function that takes an element and its index as arguments.
	ForEach(f func(T, int))

	// Join converts all elements of the slice to strings and joins them with a specified separator.
	//
	// Parameters:
	//   - sep: An optional separator string.
	//
	// Returns:
	//   A string formed by joining the string representations of the slice elements.
	Join(sep ...string) string

	// Slice returns a subset of the slice, starting at the specified begin index and optionally ending at the end index.
	//
	// Parameters:
	//   - index: A variadic parameter specifying the begin and optionally end and step indices.
	//
	// Returns:
	//   A new IAdvancedSlice[T] containing the specified subset.
	Slice(index ...int) IAdvancedSlice[T]

	// Fill sets all elements of the slice to a specified value, optionally at specified indices.
	//
	// Parameters:
	//   - value: The value to set.
	//   - indexes: An optional variadic parameter specifying the indices to fill.
	//
	// Returns:
	//   A new IAdvancedSlice[T] with the specified elements filled.
	Fill(value T, index ...int) IAdvancedSlice[T]

	// At returns the element at the specified index in the slice.
	//
	// Parameters:
	//   - index: The index of the element to retrieve.
	//
	// Returns:
	//   The element at the specified index.
	At(index int) T

	// Sort sorts the slice based on a comparison function.
	//
	// Parameters:
	//   - f: A comparison function that determines the order of elements.
	//
	// Returns:
	//   A new IAdvancedSlice[T] containing the sorted elements.
	Sort(f func(T, T) bool) IAdvancedSlice[T]

	// Values returns the underlying slice of elements.
	//
	// Returns:
	//   A slice of type []T containing all elements.
	Values() []T

	// Filter creates a new slice by filtering elements based on a predicate function.
	//
	// Parameters:
	//   - f: A predicate function that determines whether an element should be included in the new slice.
	//
	// Returns:
	//   A new slice containing elements that satisfy the predicate.
	Filter(f func(T, int) bool) []T

	// Pop removes and returns the last element from the slice.
	//
	// Returns:
	//   The last element of the slice.
	Pop() T

	// PopIs removes and returns the last element from the slice.
	//
	// Returns:
	//   The last element of the slice.
	//   A boolean indicating whether the element was successfully removed.
	PopIs() (T, bool)

	// Push adds one or more elements to the end of the slice.
	//
	// Parameters:
	//   - values: The elements to add to the slice.
	//
	// Returns:
	//   The updated slice.
	Push(values ...T) IAdvancedSlice[T]

	// PushSlice adds a slice of elements to the end of the slice.
	//
	// Parameters:
	//   - values: The slice of elements to add to the slice.
	//
	// Returns:
	//   The updated slice.
	PushSlice(values ...IAdvancedSlice[T]) IAdvancedSlice[T]

	// Shift removes and returns the first element from the slice.
	//
	// Returns:
	//   The first element of the slice.
	Shift() T

	// ShiftIs removes and returns the first element from the slice.
	//
	// Returns:
	//   The first element of the slice.
	//   A boolean indicating whether the element was successfully removed.
	ShiftIs() (T, bool)

	// Unshift adds one or more elements to the beginning of the slice.
	//
	// Parameters:
	//   - values: The elements to add to the slice.
	//
	// Returns:
	//   The updated slice.
	Unshift(values ...T) IAdvancedSlice[T]

	// UnshiftSlice adds a slice of elements to the beginning of the slice.
	//
	// Parameters:
	//   - values: The slice of elements to add to the slice.
	//
	// Returns:
	//   The updated slice.
	UnshiftSlice(values ...IAdvancedSlice[T]) IAdvancedSlice[T]

	// Reverse reverses the order of elements in the slice.
	//
	// Returns:
	//   The reversed slice.
	Reverse() IAdvancedSlice[T]

	// Remove removes elements from the slice based on a predicate function.
	//
	// Parameters:
	//   - f: A predicate function that determines whether an element should be removed.
	//
	// Returns:
	//   The updated slice.
	Remove(f func(T, int) bool) IAdvancedSlice[T]

	// RemoveAt removes an element at the specified index from the slice.
	//
	// Parameters:
	//   - index: The index of the element to remove.
	//
	// Returns:
	//   The updated slice.
	RemoveAt(index int) IAdvancedSlice[T]
}
