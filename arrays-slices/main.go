package main

import "fmt"

func main() {
	var arr [2]string

	arr[0] = "oi"
	arr[1] = "tchau"

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
