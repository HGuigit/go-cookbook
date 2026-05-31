package main

import (
	"context"
	"fmt"
	"time"
)

// Esta é a assinatura clássica do database/sql
type DB interface {
	ExecContext(ctx context.Context, query string, args ...any) error
}

// Simulador de BD
type SimuladorDB struct{}

func (s *SimuladorDB) ExecContext(ctx context.Context, query string, args ...any) error {
	select {
	case <-time.After(300 * time.Millisecond):
		fmt.Printf("Query Executada com sucesso: %s\n", query)
		return nil
	case <-ctx.Done():
		return fmt.Errorf("Erro! Query cancelada pelo BD porque o contexto expirou")
	}
}

func main() {
	fmt.Println("--- Exercício: Context e Database Timeout ---")

	banco := &SimuladorDB{}

	// TODO 1: Crie um Context com Timeout configurado para DURAR MENOS que 300ms (ex: 100ms)
	// Lembre-se de dar defer no cancel()

	// TODO 2: Chame o método ExecContext do "banco" passando o contexto criado e uma query SQL fictícia.

	// TODO 3: Verifique se houve erro no retorno e o imprima.
	// O objetivo do exercício é você provocar PROPOSITADAMENTE o erro de Timeout (Cancelamento da Query)!
}
