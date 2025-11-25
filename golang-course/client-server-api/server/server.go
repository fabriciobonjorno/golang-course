package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type ExchangeAPIResponse struct {
	USDBRL Currency `json:"USDBRL"`
}

type Currency struct {
	Bid string `json:"bid"`
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Erro ao conectar ao MySQL:", err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS cotacoes (
			id INT AUTO_INCREMENT PRIMARY KEY,
			bid VARCHAR(20),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatal("Erro ao criar tabela:", err)
	}

	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		handleCotacao(w, r, db)
	})

	fmt.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleCotacao(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	ctxAPI, cancelAPI := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancelAPI()

	req, err := http.NewRequestWithContext(ctxAPI, "GET",
		"https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		log.Println("Erro criando requisição:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Erro na chamada da API externa (timeout provável):", err)
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}
	defer resp.Body.Close()

	var result ExchangeAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Erro ao decodificar JSON:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Timeout de 10ms para escrita no banco
	ctxDB, cancelDB := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelDB()

	_, err = db.ExecContext(ctxDB, "INSERT INTO cotacoes (bid) VALUES (?)", result.USDBRL.Bid)
	if err != nil {
		log.Println("Erro ao salvar no banco (timeout provável):", err)
	}

	// Retorno JSON para o cliente
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"bid": result.USDBRL.Bid,
	})
}
