package main

import "fmt"

func easy() {
	fmt.Println("Primeiro print")
	// executa a função de fechamento no final do escopo da função easy, mesmo que ocorra um erro ou uma saída antecipada da função.
	defer fmt.Println("Ultimo print")
	fmt.Println("Segundo print")
}
