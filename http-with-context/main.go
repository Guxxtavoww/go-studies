package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

// contexto servem para controlar o tempo de vida de uma requisição HTTP. Eles permitem que você defina um prazo para a conclusão da requisição, e se esse prazo for excedido, a requisição será cancelada automaticamente. Isso é útil para evitar que a aplicação fique travada esperando uma resposta que nunca vai chegar, economizando recursos do sistema e melhorando a experiência do usuário.

// Aqui o contexto não tem timeout, então a requisição pode demorar indefinidamente para ser concluída. Isso pode ser útil em alguns casos, mas também pode levar a problemas de desempenho e travamentos se a requisição demorar muito para ser concluída.
func contextWithCancel() *context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	return &ctx
}

func main() {
	// Cria um contexto com timeout de 1 segundo. Isso significa que se a requisição demorar mais de 1 segundo, ela será cancelada automaticamente. Isso é útil para evitar que a aplicação fique travada esperando uma resposta que nunca vai chegar.
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()

	request, error := http.NewRequestWithContext(ctx, "GET", "https://www.google.com", nil)

	if error != nil {
		panic(error)
	}

	response, error := http.DefaultClient.Do(request)

	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	body, error := io.ReadAll(response.Body)

	if error != nil {
		panic(error)
	}

	println(string(body))
}
