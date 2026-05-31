package main

import (
	"context"
	"fmt"
)

// A interface simulando o serviço gerado do gRPC
type FaturamentoService interface {
	EmitirNota(ctx context.Context, req *NotaRequest) (*NotaResponse, error)
}

// TODO 1: Crie as structs "NotaRequest" (contendo Valor float64) e "NotaResponse" (contendo Codigo string)

// TODO 2: Crie uma Struct "faturamentoServer" (para simular a implementação real do serviço)

// TODO 3: Atrele um método à struct "faturamentoServer" implementando a interface FaturamentoService.
// Dentro do método, imprima o valor recebido no Request e retorne uma NotaResponse fixa simulando sucesso.

func main() {
	fmt.Println("--- Exercício: Simulação gRPC ---")

	// TODO 4: Instancie seu "faturamentoServer" e realize uma chamada direta ao método EmitirNota()
	// passando um contexto vazio (context.Background()) e imprimindo a resposta!
}
