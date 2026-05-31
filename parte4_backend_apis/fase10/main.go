package main

import (
	"context"
	"fmt"
	"log"
	"time"
	// Na vida real você importaria o driver assim:
	// _ "github.com/jackc/pgx/v5/stdlib"
)

// =====================================================================
// SIMULAÇÃO DIDÁTICA DO COMPORTAMENTO DO DATABASE/SQL
// (Usamos mocks para não exigir um container Docker de Postgres rodando)
// =====================================================================

type MockDB struct{}

// Simulando db.QueryRowContext
func (db *MockDB) QueryRowContext(ctx context.Context, query string, args ...any) (string, error) {
	fmt.Printf("[DB] Executando Query: %s com args %v\n", query, args)
	
	// Simulando um processamento demorado no banco
	select {
	case <-time.After(200 * time.Millisecond):
		return "Resultado do Banco (Sucesso)", nil
	case <-ctx.Done():
		// O banco intercepta o cancelamento do Context e interrompe a query!
		return "", fmt.Errorf("query cancelada pelo banco: %w", ctx.Err())
	}
}

func main() {
	fmt.Println("=== Fase 10: Banco de Dados, Context e Queries ===\n")
	
	// Na vida real: db, err := sql.Open("pgx", "postgres://user:pass@localhost:5432/db")
	db := &MockDB{} 

	// 1. Query Comum e Prevenção de SQL Injection
	fmt.Println("[Query Segura com Parâmetros]")
	querySegura := "SELECT nome FROM usuarios WHERE id = $1" // $1 no Postgres, ? no MySQL
	
	resultado, _ := db.QueryRowContext(context.Background(), querySegura, 42)
	fmt.Println("-> Resposta:", resultado)

	// 2. Transações e Cancelamento via Context
	fmt.Println("\n[Context Timeout no Banco de Dados]")
	// Imagine que a requisição inteira tem 100ms pra responder
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	fmt.Println("-> Disparando uma query que leva 200ms em um Context de 100ms...")
	// O QueryRowContext passa o ctx pro driver de DB. Se o tempo estourar, a query morre e libera o lock do banco.
	_, err := db.QueryRowContext(ctxTimeout, "UPDATE carteira SET saldo = saldo - 10 WHERE id = $1", 42)
	
	if err != nil {
		log.Println("Erro capturado:", err)
		fmt.Println("Ação real: Rollback da transação e retorno de HTTP 504 Gateway Timeout.")
	}
}
