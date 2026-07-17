package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func SaveToFile(data ViaCepResponse) {
	file_name := fmt.Sprintf("data-%s.json", data.Cep)
	file, err := os.Create(file_name)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		return
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
}

func main() {
	// range os.Args[1:] percorre todos os argumentos passados para o programa, exceto o primeiro (que é o nome do programa)
	for _, cep := range os.Args[1:] {
		url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

		request, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching URL: %v\n", err)
		}

		defer request.Body.Close()

		response, err := io.ReadAll(request.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading response: %v\n", err)
		}

		var data ViaCepResponse

		// não é necessario o := porque a variavel data ja foi declarada
		err = json.Unmarshal(response, &data)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error unmarshaling JSON: %v\n", err)
		}

		fmt.Println(data.Logradouro)
		SaveToFile(data)
	}
}
