package main

import (
	"context"
	"fmt"
	"log"
)

// =====================================================================
// SIMULAÇÃO DO CÓDIGO GERADO PELO PROTOC (O que o Go geraria do .proto)
// Na prática, você nunca escreve isso. Fica em um arquivo .pb.go
// =====================================================================

type UsuarioRequest struct {
	Id int32
}

type UsuarioResponse struct {
	Nome  string
	Email string
}

// A interface que o nosso servidor deverá implementar
type UsuarioServiceServer interface {
	ObterUsuario(context.Context, *UsuarioRequest) (*UsuarioResponse, error)
}

// =====================================================================
// O SEU CÓDIGO (A lógica de negócio real)
// =====================================================================

// 1. Criamos nossa Struct do Servidor
type meuServidorGRPC struct {
	// Poderia ter conexão com banco de dados aqui, logs, etc
}

// 2. Implementamos a interface exigida pelo gRPC para o método ObterUsuario
func (s *meuServidorGRPC) ObterUsuario(ctx context.Context, req *UsuarioRequest) (*UsuarioResponse, error) {
	fmt.Printf("[Servidor] Requisição gRPC recebida para ID: %d\n", req.Id)

	// Lógica de Banco de Dados simulada
	if req.Id == 1 {
		return &UsuarioResponse{
			Nome:  "Guilherme",
			Email: "guilherme@gRPC.local",
		}, nil
	}

	return nil, fmt.Errorf("usuário não encontrado")
}

func main() {
	fmt.Println("=== Fase 9: Comunicação de Alta Performance (gRPC) - Simulação ===\n")
	
	// Em um projeto real, faríamos:
	// listener, _ := net.Listen("tcp", ":50051")
	// grpcServer := grpc.NewServer()
	// pb.RegisterUsuarioServiceServer(grpcServer, &meuServidorGRPC{})
	// grpcServer.Serve(listener)
	
	fmt.Println("Instanciando o servidor e chamando os métodos diretamente (Simulando o Cliente)...\n")

	servidor := &meuServidorGRPC{}

	// Simulando uma chamada do cliente com Contexto
	req := &UsuarioRequest{Id: 1}
	
	res, err := servidor.ObterUsuario(context.Background(), req)
	if err != nil {
		log.Fatal("Erro na chamada gRPC:", err)
	}

	fmt.Println(">>> RESPOSTA DO gRPC:")
	fmt.Printf("Nome: %s | Email: %s\n", res.Nome, res.Email)
}
