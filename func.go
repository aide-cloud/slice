package slice

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func Length[T any](s []T) int {
	return len(s)
}

func Map[T, K any](list []T, f func(T, int) K) []K {
	newData := make([]K, 0, len(list))
	for index, v := range list {
		newData = append(newData, f(v, index))
	}
	return newData
}

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

func Every[T any](s []T, f func(T) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}

func Find[T any](s []T, f func(T) bool) (v T) {
	for _, item := range s {
		if f(item) {
			return item
		}
	}
	return
}

func FindIndex[T any](s []T, f func(T) bool) int {
	for i, item := range s {
		if f(item) {
			return i
		}
	}
	return -1
}

func FindLast[T any](s []T, f func(T) bool) (v T) {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return s[i]
		}
	}
	return
}

func FindLastIndex[T any](s []T, f func(T) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return i
		}
	}
	return -1
}

func ForEach[T any](s []T, f func(T, int)) {
	list := make([]T, 0, len(s))
	copy(list, s)
	for index, item := range list {
		f(item, index)
	}
}

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
		return s[begin:end:step]
	default:
		return s
	}
}

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

func String[T any](s []T) string {
	bs, err := json.Marshal(s)
	if err != nil {
		return "[]"
	}
	return string(bs)
}

func At[T any](s []T, index int) T {
	return s[index]
}

func Sort[T any](s []T, f func(a, b T) bool) []T {
	sort.Slice(s, func(i, j int) bool {
		return f(s[i], s[j])
	})
	return s
}
