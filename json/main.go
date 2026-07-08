package main

import (
	"encoding/json"
	// "os"
)

// The {} happens because json.Marshal only sees exported fields (ones starting with an uppercase letter). Your id and balance fields are lowercase, so the encoding/json package can't access them via reflection — it just serializes them as if they didn't exist.

// Fix: capitalize the field names, and optionally add JSON tags if you want the output keys to stay lowercase

type Account struct {
	Id      int
	Balance float64
}

// json:_any_' => mapeia os campos da struct para nomes de campos JSON diferentes, é possivel validar os campos do json com a tag validate, para isso é necessario importar o pacote "github.com/go-playground/validator/v10" e criar uma instância do validador, e chamar o método Struct() passando a struct que deseja validar.
type AccountWithTags struct {
	Id      int     `json:"id"`
	Balance float64 `json:"balance" validate:"required"`
}

func main() {
	account := AccountWithTags{Id: 1, Balance: 100.0}

	result, err := json.Marshal(account)

	if err != nil {
		panic(err)
	}

	println(string(result))
}
