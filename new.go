package slice

// NewAdvancedSlice 创建一个高级切片
func NewAdvancedSlice[T any](data []T) *advancedSlice[T] {
	return &advancedSlice[T]{data: data}
}
