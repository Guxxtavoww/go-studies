package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

// Exemplo de como usar Post
func usingPostRequest(client *http.Client) int64 {
	jsonData := bytes.NewBuffer([]byte(`{"name": "John", "age": 30}`))

	response, err := client.Post("https://httpbin.org/post", "application/json", jsonData)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	// io.CopyBuffer é usado para copiar o conteúdo do corpo da resposta para o stdout, usando um buffer interno. Isso é útil para evitar a alocação de memória desnecessária e melhorar a performance.
	written, err := io.CopyBuffer(os.Stdout, response.Body, nil)

	if err != nil {
		panic(err)
	}

	return written
}

func main() {
	// Tem um timeout de 1 segundo para a requisição, ou seja, se a requisição demorar mais de 1 segundo, ela será cancelada. Salvando recursivamente recursos do sistema e evitando que a aplicação fique travada esperando uma resposta que nunca vai chegar.
	client := http.Client{Timeout: time.Second}

	response, err := client.Get("https://www.google.com")

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	println(string(body))
}
