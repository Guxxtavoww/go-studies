package main

import (
	"fmt"
	"strconv"
)

func main() {
	value := "123"

	numberfyed_value, err := strconv.Atoi(value)

	if err != nil {
		panic("Invalid number")
	}

	fmt.Println(numberfyed_value)
}
