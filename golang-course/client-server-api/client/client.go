package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatal("Erro criando requisição:", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Erro ao chamar servidor:", err)
	}
	defer resp.Body.Close()

	var data map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatal("Erro ao decodificar resposta:", err)
	}

	bid := data["bid"]
	err = os.WriteFile("cotacao.txt", []byte("Dólar: "+bid), 0644)
	if err != nil {
		log.Fatal("Erro ao salvar arquivo:", err)
	}

	fmt.Println("Cotação salva em cotacao.txt")
}
