package main

import (
	"fmt"
	"log"
	"net/http"
)

// TODO 1: Crie uma função handler (ex: saudacaoHandler) que escreva "Bem vindo ao Go!" na resposta.
// Dica: A assinatura precisa ser: func(w http.ResponseWriter, r *http.Request)

func main() {
	fmt.Println("--- Exercício: Servidor HTTP ---")

	// TODO 2: Use http.HandleFunc() para associar a rota "/bemvindo" ao seu handler criado acima.

	// TODO 3: Inicie o servidor usando http.ListenAndServe() na porta ":8081".
	// (Usamos 8081 para não conflitar com a 8080 do exemplo original caso ele esteja rodando)
	
	fmt.Println("O servidor deverá subir na porta 8081. Teste no navegador: http://localhost:8081/bemvindo")
	
	// CÓDIGO AQUI:
	// log.Fatal(http.ListenAndServe(...))
}
