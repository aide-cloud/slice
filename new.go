package slice

// NewAdvancedSlice creates a new instance of an advanced slice.
//
// Parameters:
//   - data: The initial data to populate the advanced slice. This can be an empty slice or pre-populated with elements.
//
// Returns:
//
//   - IAdvancedSlice[T]: An interface of type IAdvancedSlice[T], which provides advanced functionality for manipulating the slice.
//
// Example:
//
//	// Create an advanced slice with initial data
//	advanced := NewAdvancedSlice([]int{1, 2, 3, 4})
//
//	// Create an empty advanced slice
//	emptyAdvanced := NewAdvancedSlice[int](nil)
func NewAdvancedSlice[T any](data ...T) IAdvancedSlice[T] {
	return &advancedSlice[T]{data: data}
}
