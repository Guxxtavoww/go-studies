package main

import "fmt"

type ArrType [5]int

var arr ArrType = ArrType{1, 2, 3, 4, 5}

func main() {
	// var arr [2]string

	// arr[0] = "oi"
	// arr[1] = "tchau"

	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	secound_array := []string{"a", "b", "c"}

	for i := 0; i < len(secound_array); i++ {
		fmt.Println(secound_array[i])
	}

	for index, num := range arr {
		fmt.Printf("O valor do indice %d é %d", index, num)
	}
}
