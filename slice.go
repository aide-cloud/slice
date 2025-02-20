package slice

var _ IAdvancedSlice[any] = (*advancedSlice[any])(nil)

// advancedSlice 高级切片
type advancedSlice[T any] struct {
	data []T
}

func (s *advancedSlice[T]) String() string {
	return String(s.data)
}

func (s *advancedSlice[T]) Map(f func(T, int) T) IAdvancedSlice[T] {
	s.data = Map(s.data, f)
	return s
}

func (s *advancedSlice[T]) Unique(f func(T) string) IAdvancedSlice[T] {
	s.data = Unique(s.data, f)
	return s
}

func (s *advancedSlice[T]) Concat(ss ...advancedSlice[T]) IAdvancedSlice[T] {
	list := make([][]T, 0, len(ss))
	list = append(list, s.data)
	for _, a := range ss {
		list = append(list, a.data)
	}
	s.data = Concat(list...)
	return s
}

func (s *advancedSlice[T]) CopyWithIn(indexes ...int) IAdvancedSlice[T] {
	s.data = CopyWithIn(s.data, indexes...)
	return s
}

func (s *advancedSlice[T]) Every(f func(T) bool) bool {
	return Every(s.data, f)
}

func (s *advancedSlice[T]) Find(f func(T) bool) T {
	return Find(s.data, f)
}

func (s *advancedSlice[T]) FindIndex(f func(T) bool) int {
	return FindIndex(s.data, f)
}

func (s *advancedSlice[T]) FindLast(f func(T) bool) T {
	return FindLast(s.data, f)
}

func (s *advancedSlice[T]) FindLastIndex(f func(T) bool) int {
	return FindLastIndex(s.data, f)
}

func (s *advancedSlice[T]) ForEach(f func(T, int)) {
	ForEach(s.data, f)
}

func (s *advancedSlice[T]) Join(sep ...string) string {
	return Join(s.data, sep...)
}

func (s *advancedSlice[T]) Slice(index ...int) IAdvancedSlice[T] {
	s.data = Slice(s.data, index...)
	return s
}

func (s *advancedSlice[T]) Fill(value T, index ...int) IAdvancedSlice[T] {
	s.data = Fill(s.data, value, index...)
	return s
}

func (s *advancedSlice[T]) At(index int) T {
	return At(s.data, index)
}

func (s *advancedSlice[T]) Sort(f func(T, T) bool) IAdvancedSlice[T] {
	s.data = Sort(s.data, f)
	return s
}

func (s *advancedSlice[T]) Values() []T {
	return s.data
}

// Length 返回切片的长度
func (s *advancedSlice[T]) Length() int {
	return Length(s.data)
}
