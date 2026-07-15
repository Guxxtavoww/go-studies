package main

import (
	"fmt"
	"log"
	"net/http"
)

// http.HandleFunc registra a rota num mux global e implícito (DefaultServeMux), enquanto http.NewServeMux() cria um mux próprio e explícito, dando controle isolado sobre as rotas do servidor.

var SERVER_PORT string = ":8080"

// usando mux é possivel ter multiplos servidores rodando na mesma porta ou em portas diferentes para propositos diferentes, cada um com suas rotas e handlers, sem conflitar entre si.

func main() {
	mux_server := http.NewServeMux()

	mux_server.HandleFunc("/", handler)
	mux_server.Handle("/blog", &Blog{title: "My Blog"})

	fmt.Println("Server is running on http://localhost" + SERVER_PORT)

	err := http.ListenAndServe(SERVER_PORT, mux_server)

	if err != nil {
		log.Fatal(err)
	}
}

func handler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Hello, World!"))
}

type Blog struct {
	title string
}

// ServeHTTP é o método que implementa a interface http.Handler, permitindo que a struct Blog seja usada como um manipulador de requisições HTTP.
func (b *Blog) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/blog" {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(b.title))
}
