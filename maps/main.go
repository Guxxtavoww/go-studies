package main

import "fmt"

func main() {
	// O object do GoLang
	salaries := map[string]int{
		"gustavo": 2,
		"kleber":  20,
		"maria":   1002,
	}

	fmt.Println(salaries["gustavo"])

	// Deleta com base chave do map
	delete(salaries, "gustavo")

	_, exists := salaries["gustavo"]

	if !exists {
		fmt.Println("Funcionou")
	}

	map_without_values := make(map[string]int)

	map_without_values["Fodase"] = 232132132132

	fmt.Println(map_without_values["Fodase"])

	// O equivalente de Object.entries
	for key, value := range map_without_values {
		fmt.Printf("O salario de %s é %d \n", key, value)
	}
}
