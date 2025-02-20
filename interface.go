package slice

import (
	"fmt"
)

// IAdvancedSlice advanced slice interface
type IAdvancedSlice[T any] interface {
	fmt.Stringer

	// Length get length of slice
	Length() int

	Map(f func(T, int) T) IAdvancedSlice[T]

	Unique(f func(T) string) IAdvancedSlice[T]

	Concat(s ...advancedSlice[T]) IAdvancedSlice[T]

	CopyWithIn(indexes ...int) IAdvancedSlice[T]

	Every(f func(T) bool) bool

	Find(f func(T) bool) T

	FindIndex(f func(T) bool) int

	FindLast(f func(T) bool) T

	FindLastIndex(f func(T) bool) int

	ForEach(f func(T, int))

	Join(sep ...string) string

	Slice(index ...int) IAdvancedSlice[T]

	Fill(value T, index ...int) IAdvancedSlice[T]

	At(index int) T

	Sort(f func(T, T) bool) IAdvancedSlice[T]

	Values() []T
}
