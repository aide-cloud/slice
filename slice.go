package slice

var _ IAdvancedSlice[any] = (*advancedSlice[any])(nil)

// advancedSlice is a concrete implementation of the IAdvancedSlice interface.
// It provides advanced functionality for manipulating slices.
type advancedSlice[T any] struct {
	data []T
}

// String returns a string representation of the slice.
//
// Returns:
//
//   - string: A JSON string representation of the slice, or "[]" if conversion fails.
func (s *advancedSlice[T]) String() string {
	return String(s.data)
}

// Map applies a transformation function to each element of the slice and returns a new advanced slice with the transformed elements.
//
// Parameters:
//   - f: A function that takes an element of type T and its index, and returns an element of type T.
//
// Returns:
//
//   - IAdvancedSlice[T]: A new IAdvancedSlice[T] containing the results of applying the transformation function to each element.
func (s *advancedSlice[T]) Map(f func(T, int) T) IAdvancedSlice[T] {
	s.data = Map(s.data, f)
	return s
}

// Unique returns a new advanced slice with unique elements based on a key function.
//
// Parameters:
//
//   - f: A function that extracts a key from each element of type T. The key must be a string.
//
// Returns:
//
//   - IAdvancedSlice[T]: A new IAdvancedSlice[T] containing only the first occurrence of each unique key.
func (s *advancedSlice[T]) Unique(f func(T) string) IAdvancedSlice[T] {
	s.data = Unique(s.data, f)
	return s
}

// Concat concatenates multiple slices into one and returns a new advanced slice.
//
// Parameters:
//   - ss: A variadic parameter representing multiple slices to concatenate.
//
// Returns:
//
//   - IAdvancedSlice[T]: A new IAdvancedSlice[T] containing all elements from the input slices.
func (s *advancedSlice[T]) Concat(ss ...IAdvancedSlice[T]) IAdvancedSlice[T] {
	list := make([][]T, 0, len(ss))
	list = append(list, s.data)
	for _, a := range ss {
		list = append(list, a.Values())
	}
	s.data = Concat(list...)
	return s
}

// CopyWithIn creates a new advanced slice containing elements at specified indices from the original slice.
//
// Parameters:
//
//   - indexes: A variadic parameter representing the indices of elements to include in the new slice.
//
// Returns:
//
//   - IAdvancedSlice[T]: A new IAdvancedSlice[T] containing elements at the specified indices.
func (s *advancedSlice[T]) CopyWithIn(indexes ...int) IAdvancedSlice[T] {
	s.data = CopyWithIn(s.data, indexes...)
	return s
}

// Every check if all elements in the slice satisfy a given predicate function.
//
// Parameters:
//
//   - f: A predicate function that takes an element and returns a boolean.
//
// Returns:
//
//   - bool: true if all elements satisfy the predicate, false otherwise.
func (s *advancedSlice[T]) Every(f func(T) bool) bool {
	return Every(s.data, f)
}

// Find searches for the first element in the slice that satisfies a given predicate function.
//
// Parameters:
//
//   - f: A predicate function that takes an element and returns a boolean.
//
// Returns:
//
//   - T: The first element that satisfies the predicate, or the zero value if no such element exists.
func (s *advancedSlice[T]) Find(f func(T) bool) T {
	return Find(s.data, f)
}

// FindIndex finds the index of the first element in the slice that satisfies a given predicate function.
//
// Parameters:
//
//   - f: A predicate function that takes an element and returns a boolean.
//
// Returns:
//
//   - int (index): The index of the first element that satisfies the predicate, or -1 if no such element exists.
func (s *advancedSlice[T]) FindIndex(f func(T) bool) int {
	return FindIndex(s.data, f)
}

// FindLast searches for the last element in the slice that satisfies a given predicate function.
//
// Parameters:
//
//   - f: A predicate function that takes an element and returns a boolean.
//
// Returns:
//
//   - T: The last element that satisfies the predicate, or the zero value if no such element exists.
func (s *advancedSlice[T]) FindLast(f func(T) bool) T {
	return FindLast(s.data, f)
}

// FindLastIndex finds the index of the last element in the slice that satisfies a given predicate function.
//
// Parameters:
//
//   - f: A predicate function that takes an element and returns a boolean.
//
// Returns:
//
//   - int (index): The index of the last element that satisfies the predicate, or -1 if no such element exists.
func (s *advancedSlice[T]) FindLastIndex(f func(T) bool) int {
	return FindLastIndex(s.data, f)
}

// ForEach iterates over each element in the slice and applies a provided function to it.
//
// Parameters:
//
//   - f: A function that takes an element and its index as arguments.
func (s *advancedSlice[T]) ForEach(f func(T, int)) {
	ForEach(s.data, f)
}

// Join converts all elements of the slice to strings and joins them with a specified separator.
//
// Parameters:
//
//   - sep: An optional separator string.
//
// Returns:
//
//   - string: A string formed by joining the string representations of the slice elements.
func (s *advancedSlice[T]) Join(sep ...string) string {
	return Join(s.data, sep...)
}

// Slice returns a subset of the slice, starting at the specified begin index and optionally ending at the end index.
//
// Parameters:
//
//   - index: A variadic parameter specifying the begin and optionally end and step indices.
//
// Returns:
//
//   - IAdvancedSlice[T]: A new IAdvancedSlice[T] containing the specified subset.
func (s *advancedSlice[T]) Slice(index ...int) IAdvancedSlice[T] {
	s.data = Slice(s.data, index...)
	return s
}

// Fill sets all elements of the slice to a specified value, optionally at specified indices.
//
// Parameters:
//
//   - value: The value to set.
//   - indexes: An optional variadic parameter specifying the indices to fill.
//
// Returns:
//
//   - IAdvancedSlice[T]: A new IAdvancedSlice[T] with the specified elements filled.
func (s *advancedSlice[T]) Fill(value T, index ...int) IAdvancedSlice[T] {
	s.data = Fill(s.data, value, index...)
	return s
}

// At returns the element at the specified index in the slice.
//
// Parameters:
//
//   - index: The index of the element to retrieve.
//
// Returns:
//
//   - T: The element at the specified index.
func (s *advancedSlice[T]) At(index int) T {
	return At(s.data, index)
}

// Sort sorts the slice based on a comparison function.
//
// Parameters:
//
//   - f: A comparison function that determines the order of elements.
//
// Returns:
//
//   - IAdvancedSlice[T]: A new IAdvancedSlice[T] containing the sorted elements.
func (s *advancedSlice[T]) Sort(f func(T, T) bool) IAdvancedSlice[T] {
	s.data = Sort(s.data, f)
	return s
}

// Values returns the underlying slice of elements.
//
// Returns:
//
//   - []T: A slice of type []T containing all elements.
func (s *advancedSlice[T]) Values() []T {
	return s.data
}

// Length returns the number of elements in the slice.
//
// Returns:
//
//   - int (len): The length of the slice as an integer.
func (s *advancedSlice[T]) Length() int {
	return Length(s.data)
}

// Filter creates a new slice containing elements that satisfy a given predicate function.
//
// Parameters:
//
//   - f: A predicate function that takes an element and its index as arguments and returns a boolean.
//
// Returns:
//
//   - []T: A new IAdvancedSlice[T] containing the filtered elements.
func (s *advancedSlice[T]) Filter(f func(T, int) bool) []T {
	return Filter(s.data, f)
}

// Pop removes and returns the last element from the slice.
//
// Returns:
//
//   - T: The last element of the slice.
func (s *advancedSlice[T]) Pop() T {
	length := s.Length()
	if length == 0 {
		var zero T
		return zero
	}

	last := s.At(length - 1)
	s.data = s.data[:length-1]
	return last
}

// PopIs removes and returns the last element from the slice.
//
// Returns:
//
//   - T: The last element of the slice.
//   - bool: A boolean indicating whether the operation was successful.
func (s *advancedSlice[T]) PopIs() (T, bool) {
	length := s.Length()
	if length == 0 {
		var zero T
		return zero, false
	}

	last := s.At(length - 1)
	s.data = s.data[:length-1]
	return last, true
}

// Push adds one or more elements to the end of the slice.
//
// Parameters:
//
//   - values: One or more elements to add.
//
// Returns:
//
//   - IAdvancedSlice[T]: A new IAdvancedSlice[T] with the added elements.
func (s *advancedSlice[T]) Push(values ...T) IAdvancedSlice[T] {
	s.data = append(s.data, values...)
	return s
}

// PushSlice adds one or more slices to the end of the slice.
//
// Parameters:
//
//   - values: One or more slices to add.
//
// Returns:
//
//   - IAdvancedSlice[T]： A new IAdvancedSlice[T] with the added slices.
func (s *advancedSlice[T]) PushSlice(values ...IAdvancedSlice[T]) IAdvancedSlice[T] {
	for _, v := range values {
		s.data = append(s.data, v.Values()...)
	}
	return s
}

// Shift removes and returns the first element from the slice.
//
// Returns:
//
//   - T: The first element of the slice.
func (s *advancedSlice[T]) Shift() T {
	length := s.Length()
	if length == 0 {
		var zero T
		return zero
	}

	first := s.At(0)
	s.data = s.data[1:]
	return first
}

// ShiftIs removes and returns the first element from the slice.
//
// Returns:
//
//   - T: The first element of the slice.
//   - bool: A boolean indicating whether the operation was successful.
func (s *advancedSlice[T]) ShiftIs() (T, bool) {
	length := s.Length()
	if length == 0 {
		var zero T
		return zero, false
	}

	first := s.At(0)
	s.data = s.data[1:]
	return first, true
}

// Unshift adds one or more elements to the beginning of the slice.
//
// Parameters:
//
//   - values: One or more elements to add.
//
// Returns:
//
//   - IAdvancedSlice[T]: A new IAdvancedSlice[T] with the added elements.
func (s *advancedSlice[T]) Unshift(values ...T) IAdvancedSlice[T] {
	s.data = append(values, s.data...)
	return s
}

// UnshiftSlice adds one or more slices to the beginning of the slice.
//
// Parameters:
//
//   - values: One or more slices to add.
//
// Returns:
//
//   - IAdvancedSlice[T]: A new IAdvancedSlice[T] with the added slices.
func (s *advancedSlice[T]) UnshiftSlice(values ...IAdvancedSlice[T]) IAdvancedSlice[T] {
	for _, v := range values {
		s.data = append(v.Values(), s.data...)
	}
	return s
}

// Reverse reverses the order of elements in the slice.
//
// Returns:
//
//   - IAdvancedSlice[T]: A new IAdvancedSlice[T] with reversed elements.
func (s *advancedSlice[T]) Reverse() IAdvancedSlice[T] {
	s.data = Reverse(s.data)
	return s
}

// Remove removes elements from the slice based on a predicate function.
//
// Parameters:
//   - f: A predicate function that takes an element and its index as arguments and returns a boolean.
//
// Returns:
//
//   - IAdvancedSlice[T]: A new IAdvancedSlice[T] with the removed elements.
func (s *advancedSlice[T]) Remove(f func(T, int) bool) IAdvancedSlice[T] {
	s.data = Remove(s.data, f)
	return s
}

// RemoveAt removes an element at the specified index from the slice.
//
// Parameters:
//
//   - index: The index of the element to remove.
//
// Returns:
//
//   - IAdvancedSlice[T]: A new IAdvancedSlice[T] with the removed element.
func (s *advancedSlice[T]) RemoveAt(index int) IAdvancedSlice[T] {
	s.data = RemoveAt(s.data, index)
	return s
}
