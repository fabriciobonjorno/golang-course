package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type AddressResult struct {
	Api      string `json:"api"`
	Response map[string]interface{}
	Err      error
}

func fetch(ctx context.Context, url, apiLink string, ch chan<- AddressResult) {
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ch <- AddressResult{Api: apiLink, Err: err}
		return
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		ch <- AddressResult{Api: apiLink, Err: err}
	}

	ch <- AddressResult{Api: apiLink, Err: err}
}

func main() {

	cep := "78150190"
	ch := make(chan AddressResult, 2)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// URLs
	urlBrasilAPI := "https://brasilapi.com.br/api/cep/v1/" + cep
	urlViaCEP := "http://viacep.com.br/ws/" + cep + "/json/"

	go fetch(ctx, urlBrasilAPI, "BrasilAPI", ch)
	go fetch(ctx, urlViaCEP, "ViaCEP", ch)

	select {
	case result := <-ch:
		if result.Err != nil {
			fmt.Println("Erro:", result.Err)
			return
		}

		fmt.Println("🔥 API mais rápida:", result.Api)
		fmt.Println("📍 Endereço recebido:")
		for k, v := range result.Response {
			fmt.Printf("%s: %v\n", k, v)
		}

	case <-ctx.Done():
		fmt.Println("❌ Timeout: Nenhuma API respondeu dentro de 1 segundo.")
	}
}
