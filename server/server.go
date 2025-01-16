package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Cotacao struct {
	Bid string `json:"bid"`
}

func main() {

	http.HandleFunc("/cotacao", HandleCotacao)
	log.Println("Servidor iniciado na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HandleCotacao(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	defer log.Println("Request completed")

	// consume api with context
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		log.Printf("Erro ao criar requisição para API: %v", err)
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Erro ao chamar API: %v", err)
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Printf("Resposta inesperada da API: %s", res.Status)
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}

	var result map[string]Cotacao
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Printf("Erro ao decodificar resposta da API: %v", err)
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}

	cotacao := result["USDBRL"]
	println(cotacao.Bid)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(cotacao.Bid))

}
