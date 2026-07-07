package main

import (
	"io"
	"net/http"
)

func main() {
	request, err := http.Get("https://www.google.com")

	if err != nil {
		panic(err)
	}

	response, error := io.ReadAll(request.Body)

	if error != nil {
		panic(error)
	}

	println(string(response))

	// É necessario fechar o corpo da resposta para liberar recursos
	request.Body.Close()
}
