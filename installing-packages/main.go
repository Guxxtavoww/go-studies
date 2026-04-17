package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Print[T constraints.Signed](v T) {
	fmt.Println(v)
}

func main() {
	Print(42)       // int
	Print(int8(10)) // int8
	Print(int64(9)) // int64
}
