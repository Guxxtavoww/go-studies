package main

import "fmt"

// A sintaxe array[start:end:max] fatia o slice: __start__ onde o slice começa, end a length formula -> len = end - start, max o capacity cap = max - start
func main() {
	// arr[:] -> tudo [1,2,3,8,4]
	number_array := []int{1, 2, 3, 8, 4}

	// cap = tamanho_original
	fmt.Printf("length=%d capacity=%d %v\n", len(number_array), cap(number_array), number_array)

	// arr[:0] do início até índice 0 (exclusivo) -> []
	fmt.Printf("length=%d capacity=%d %v\n", len(number_array[:0]), cap(number_array[:0]), number_array[:0])

	// arr[:4]do início até índice 4 (exclusivo) -> [1,2,3,8]
	fmt.Printf("length=%d capacity=%d %v\n", len(number_array[:4]), cap(number_array[:4]), number_array[:4])

	//arr[2:] do índice 2 até o fim -> [3,8,4]
	fmt.Printf("length=%d capacity=%d %v\n", len(number_array[2:]), cap(number_array[:2]), number_array[:2])

	number_array = append(number_array, 12)

	// EX: arr[1:3] do índice 1 até 3 (exclusivo) [2,3]

	fmt.Printf("length=%d capacity=%d %v\n", len(number_array[2:]), cap(number_array[:2]), number_array[:2])
}
