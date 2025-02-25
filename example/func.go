package main

import (
	"fmt"

	"github.com/aide-cloud/slice"
)

func main() {
	s := slice.NewAdvancedSlice[any](1, 2, "a")
	fmt.Println(s.Length())
}
