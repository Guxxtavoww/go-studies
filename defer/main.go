package main

import (
	"io"
	"net/http"
)

func main() {
	request, err := http.Get("https://dummyjson.com/test")

	if err != nil {
		panic(err)
	}

	response, err := io.ReadAll(request.Body)

	if err != nil {
		panic(err)
	}

	println(string(response))

	// defer serve para garantir que o recurso seja fechado mesmo que ocorra um erro ou uma saída antecipada da função. Atrasando a execução da função de fechamento até o final do escopo da função main.
	defer request.Body.Close()
}
