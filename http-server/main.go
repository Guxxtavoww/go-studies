package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

// Retorna um ponteiro para viaCepResponse porque a struct ViaCepResponse é relativamente grande e complexa, e retornar um ponteiro evita cópias desnecessárias de memória. Além disso, ao retornar um ponteiro, podemos modificar os dados da struct diretamente se necessário, sem precisar criar uma cópia da struct inteira.
func FechViaCepData(cep string) (*ViaCepResponse, error) {
	viaCepUrl := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	response, err := http.Get(viaCepUrl)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var data ViaCepResponse

	// essa linha é responsável por decodificar o corpo da resposta HTTP (que está em formato JSON) e armazenar os dados na struct data do tipo ViaCepResponse. A função json.Unmarshal recebe dois argumentos: o primeiro é o slice de bytes body, que contém os dados JSON, e o segundo é um ponteiro para a struct data, onde os dados decodificados serão armazenados. Se a decodificação for bem-sucedida, os campos da struct data serão preenchidos com os valores correspondentes do JSON.
	err = json.Unmarshal(body, &data)

	return &data, err
}

// Request é um ponteiro porque o http.Request é uma struct grande e complexa, e passar um ponteiro evita cópias desnecessárias de memória. Além disso, o http.Request contém informações sobre a requisição HTTP, como cabeçalhos, parâmetros de URL, corpo da requisição, etc. Ao usar um ponteiro, podemos acessar e modificar essas informações diretamente sem precisar criar uma cópia da struct inteira.
func SearchCepHandler(response_writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		response_writer.WriteHeader(http.StatusNotFound)

		return
	}

	response_writer.Header().Set("Content-Type", "application/json")
	response_writer.WriteHeader(http.StatusOK)

	cepQuery := request.URL.Query().Get("cep")

	if cepQuery == "" {
		response_writer.WriteHeader(http.StatusBadRequest)
		response_writer.Write([]byte(`{"error": "CEP não fornecido"}`))
	}

	viaCepResponse, err := FechViaCepData(cepQuery)

	if err != nil {
		response_writer.WriteHeader(http.StatusInternalServerError)
		response_writer.Write([]byte(`{"error": "Erro ao buscar dados do CEP"}`))
	}

	// essa linha de código é responsável por codificar a struct viaCepResponse em formato JSON e escrever essa resposta no corpo da resposta HTTP. A função json.NewEncoder(response_writer) cria um novo codificador JSON que escreve diretamente no response_writer, que é o objeto responsável por enviar a resposta ao cliente. Em seguida, o método Encode(viaCepResponse) é chamado para converter a struct viaCepResponse em JSON e enviá-la como resposta ao cliente.
	json.NewEncoder(response_writer).Encode(viaCepResponse)
}

func main() {
	http.HandleFunc("/", SearchCepHandler)

	http.ListenAndServe(":8080", nil)
}
