package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Address struct {
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

func main() {
	for _, zipCode := range os.Args[1:] {
		// req, err := http.Get(url)
		req, err := http.Get("https://viacep.com.br/ws/" + zipCode + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Request error: %v\n", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading response: %v\n", err)
		}

		var data Address
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error unmarshaling JSON: %v\n", err)
		}

		file, err := os.Create("address.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		}

		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("CEP: %s\nLogradouro: %s\nBairro: %s\nLocalidade: %s\nUF: %s\n", data.Cep, data.Logradouro, data.Bairro, data.Localidade, data.Uf))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
		}

		fmt.Printf("CEP: %s\nLogradouro: %s\nBairro: %s\nLocalidade: %s\nUF: %s\n", data.Cep, data.Logradouro, data.Bairro, data.Localidade, data.Uf)
	}
}
