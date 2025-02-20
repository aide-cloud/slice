## Slice Utility Project README

### Overview
The Slice Utility Project provides a comprehensive set of utility functions to manipulate slices in Go. These utilities are designed to simplify common operations such as filtering, mapping, and reducing slices, thereby enhancing productivity and code readability.

### Features
- **Filter**: Remove elements from a slice based on a condition.
- **Map**: Transform elements in a slice using a provided function.
- **Reduce**: Reduce a slice to a single value by applying a function cumulatively to the elements.
- **Unique**: Return a new slice with unique elements based on a key function.
- **Concat**: Concatenate multiple slices into one.
- **CopyWithIn**: Create a new slice containing elements at specified indices.
- **Every**: Check if all elements satisfy a given predicate.
- **Find**: Search for the first element that satisfies a given predicate.
- **FindIndex**: Find the index of the first element that satisfies a given predicate.
- **FindLast**: Search for the last element that satisfies a given predicate.
- **FindLastIndex**: Find the index of the last element that satisfies a given predicate.
- **ForEach**: Iterate over each element and apply a provided function.
- **Join**: Convert all elements to strings and join them with a specified separator.
- **Slice**: Return a subset of the slice.
- **Fill**: Set all elements of the slice to a specified value.
- **At**: Retrieve the element at a specified index.
- **Sort**: Sort the slice based on a comparison function.
- **Values**: Return the underlying slice of elements.

### Installation

To install the Slice Utility Project, use the following command:

```go
go get github.com/aide-cloud/slice@latest
```

To use the Slice Utility Project in your Go project, you can import it as follows:

```go
import "github.com/aide-cloud/slice"
```

### Usage Examples

#### Creating an Advanced Slice
You can create an advanced slice with initial data or an empty advanced slice.

```go
// Create an advanced slice with initial data
advanced := NewAdvancedSlice([]int{1, 2, 3, 4})

// Create an empty advanced slice
emptyAdvanced := NewAdvancedSlice[int](nil)
```


#### Applying Map Function
Transform elements in the slice using a provided function.

```go
transformed := advanced.Map(func(item int, index int) int {
    return item * 2
})
```


#### Filtering Elements
Remove elements from a slice based on a condition.

```go
filtered := advanced.Filter(func(item int, index int)) bool {
    return item > 2
})
```


#### Finding Elements
Search for elements that satisfy a given predicate.

```go
firstMatch := advanced.Find(func(item int) bool {
    return item == 3
})

index := advanced.FindIndex(func(item int) bool {
    return item == 3
})
```


#### Sorting Elements
Sort the slice based on a comparison function.

```go
sorted := advanced.Sort(func(a int, b int) bool {
    return a < b
})
```


#### Concatenating Slices
Concatenate multiple slices into one.

```go
concatenated := advanced.Concat(NewAdvancedSlice([]int{5, 6, 7}))
```


#### Unique Elements
Return a new slice with unique elements based on a key function.

```go
unique := advanced.Unique(func(item int) string {
    return fmt.Sprintf("%d", item)
})
```


#### Iterating Over Elements
Iterate over each element and apply a provided function.

```go
advanced.ForEach(func(item int, index int) {
    fmt.Printf("Element %d: %d\n", index, item)
})
```


This README provides an overview of the Slice Utility Project along with examples of how to use its various functionalities. For more detailed information, please refer to the inline documentation within the source code.