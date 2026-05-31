package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Estrutura do Payload
type Produto struct {
	ID    int     `json:"id"`
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}

// 1. Handler Básico (GET)
func pingHandler(w http.ResponseWriter, r *http.Request) {
	// Apenas para exemplificar que podemos restringir verbos manualmente se não usarmos o Mux novo
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"mensagem": "pong"}`))
}

// 2. Handler de API (POST com JSON)
func criarProdutoHandler(w http.ResponseWriter, r *http.Request) {
	var p Produto

	// Desserialização: Lendo do Body da Request (r.Body) para a struct 'p'
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Erro ao processar JSON", http.StatusBadRequest)
		return
	}

	// Lógica fake de negócio: salvar banco e gerar ID
	p.ID = 99

	// Serialização: Escrevendo a resposta com status 201 (Created)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	
	// Transforma a struct 'p' em JSON e escreve na Response (w)
	json.NewEncoder(w).Encode(p)
}

// 3. Padrão Middleware (Log e Tempo de Execução)
func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inicio := time.Now()

		// Executa o próximo handler na corrente
		next.ServeHTTP(w, r)

		duracao := time.Since(inicio)
		log.Printf("[%s] %s - demorou %v\n", r.Method, r.URL.Path, duracao)
	})
}

func main() {
	// Roteador Padrão
	mux := http.NewServeMux()

	// Definindo as rotas
	// Se usar Go 1.22+, a sintaxe "POST /api/produtos" e "GET /ping" é nativamente suportada
	mux.HandleFunc("/ping", pingHandler)
	mux.HandleFunc("/api/produtos", criarProdutoHandler)

	// Envolvendo o Mux com o Middleware
	handlerFinal := loggerMiddleware(mux)

	porta := ":8080"
	fmt.Printf("=== Servidor Iniciado na porta %s ===\n", porta)
	fmt.Println("Testes:")
	fmt.Println("curl -X GET http://localhost:8080/ping")
	fmt.Println(`curl -X POST http://localhost:8080/api/produtos -d '{"nome":"Mouse", "preco":150.00}'`)
	
	// Inicializando Servidor com Segurança (Timeout previne Slowloris)
	server := &http.Server{
		Addr:         porta,
		Handler:      handlerFinal,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}
