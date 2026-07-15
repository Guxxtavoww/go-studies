package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// cria um servidor de arquivos que serve arquivos do diretório "./public"
	file_server := http.FileServer(http.Dir("./public"))

	// adiciona o servidor de arquivos ao mux para lidar com todas as requisições na raiz "/"
	mux.Handle("/", file_server)

	// é possivel mesclar o servidor de arquivos com outros handlers, como por exemplo:
	mux.HandleFunc("/blog", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Bem-vindo ao blog!"))
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
