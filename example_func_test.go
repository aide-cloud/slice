package slice_test

import (
	"fmt"

	"github.com/aide-cloud/slice"
)

func ExampleLength() {
	fmt.Println(slice.Length([]int{1, 2, 3}))
	fmt.Println(slice.Length([]string{"a", "b", "c"}))
	fmt.Println(slice.Length([]struct{}{}))
	fmt.Println(slice.Length([]map[string]int{}))
	fmt.Println(slice.Length([]interface{}{}))
	// Output:
	// 3
	// 3
	// 0
	// 0
	// 0
}
