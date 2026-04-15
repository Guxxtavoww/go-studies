package main

import "fmt"

// Interface vazia implementa todas as structs o equivalente de any
type x interface{}

func show_type(t x) {
	fmt.Printf("O tipo é: %T e o valor é %v\n", t, t)
}

func main() {
	value := 10

	show_type(value)
}
