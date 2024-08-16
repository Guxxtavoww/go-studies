package main

import "fmt"

func main() {
	fmt.Println("Insira uma string qualquer: ")
	var user_input string

	fmt.Scan(&user_input)

	fmt.Println("String: ", user_input)
}
