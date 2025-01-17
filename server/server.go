package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func fetchCotacao(ctx context.Context) (Cotacao, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return Cotacao{}, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Cotacao{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return Cotacao{}, err
	}

	var result map[string]Cotacao
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return Cotacao{}, err
	}

	return result["USDBRL"], nil
}

func saveCotacao(ctx context.Context, cotacao Cotacao) error {
	db, err := sql.Open("sqlite3", "./cotacao.db")
	if err != nil {
		return err
	}
	defer db.Close()

	createTableQuery := `CREATE TABLE IF NOT EXISTS cotacoes (id INTEGER PRIMARY KEY AUTOINCREMENT, bid TEXT, timestamp DATETIME DEFAULT CURRENT_TIMESTAMP)`
	if _, err := db.ExecContext(ctx, createTableQuery); err != nil {
		return err
	}

	insertQuery := `INSERT INTO cotacoes (bid) VALUES (?)`
	if _, err := db.ExecContext(ctx, insertQuery, cotacao.Bid); err != nil {
		return err
	}

	return nil
}

type Cotacao struct {
	Bid string `json:"bid"`
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", HandleCotacao)
	log.Println("Servidor iniciado na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func HandleCotacao(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	defer log.Println("Request completed")

	// consume api with context
	cotacao, err := fetchCotacao(ctx)
	if err != nil {
		log.Printf("Erro ao chamar API: %v", err)
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}
	ctxBD, cancelBD := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelBD()
	err = saveCotacao(ctxBD, cotacao)
	if err != nil {
		log.Printf("Erro ao salvar no banco de dados: %v", err)
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}
	println(cotacao.Bid)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"bid": cotacao.Bid}); err != nil {
		log.Printf("Erro ao enviar resposta: %v", err)
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}
}
