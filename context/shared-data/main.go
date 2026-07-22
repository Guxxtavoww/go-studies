package main

import (
	"context"
	"fmt"
)

// 1. Define uma chave customizada (nunca use string crua)
type ContextKey string

const userIdKey ContextKey = "user_id"

func extract_user_id_from_context(ctx context.Context) string {
	// acha o valor da chave dentro do contexto e faz um parse para string
	user_id, ok := ctx.Value(userIdKey).(string)

	if !ok {
		panic("não há user_id no contexto")
	}

	return user_id
}

func wrapper_fn(ctx context.Context) {
	fmt.Println("Wrapper running")

	user_id := extract_user_id_from_context(ctx)

	fmt.Printf("id: %v", user_id)
}

func main() {
	ctx := context.Background()

	fmt.Println("Insira o id do usuario: ")

	var user_id string

	fmt.Scan(&user_id)

	if len(user_id) == 0 {
		panic("Id de usuario invalid")
	}

	// adiciona o valor do id do usuario no contexto
	ctx = context.WithValue(ctx, userIdKey, user_id)

	wrapper_fn(ctx)
}
