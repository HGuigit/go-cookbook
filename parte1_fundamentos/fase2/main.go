package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== Fase 2: Estruturas de Controle de Fluxo ===\n")

	// 1. if e else tradicional e o Short Statement
	fmt.Println("[if/else e Short Statement]")
	
	// A variável 'pontuacao' só existe dentro deste bloco if-else
	if pontuacao := rand.Intn(100); pontuacao >= 70 {
		fmt.Printf("Aprovado com %d pontos.\n", pontuacao)
	} else {
		fmt.Printf("Reprovado com %d pontos.\n", pontuacao)
	}
	// fmt.Println(pontuacao) -> GERARIA ERRO DE COMPILAÇÃO (variável fora de escopo)

	// 2. O Laço 'for' (O único do Go)
	fmt.Println("\n[Laços 'for']")

	// a. for Tradicional
	fmt.Print("Tradicional: ")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// b. for como "while"
	fmt.Print("Como 'while': ")
	contador := 3
	for contador > 0 {
		fmt.Printf("%d ", contador)
		contador--
	}
	fmt.Println()

	// c. Iterando com 'range' sobre um Slice
	nomes := []string{"Alice", "Bob", "Carlos"}
	fmt.Println("\n[Iteração com 'range']")
	for indice, nome := range nomes {
		fmt.Printf("Índice: %d, Nome: %s\n", indice, nome)
	}

	// Omitindo o índice com Blank Identifier (_)
	for _, nome := range nomes {
		fmt.Printf("Nome sem índice: %s\n", nome)
	}

	// 3. O switch modernizado
	fmt.Println("\n[O Switch]")
	dia := time.Tuesday

	switch dia {
	case time.Monday:
		fmt.Println("Segunda-feira. Início de semana.")
	case time.Tuesday, time.Wednesday, time.Thursday:
		fmt.Println("Meio de semana de trabalho.")
	case time.Friday:
		fmt.Println("Sexta-feira. Quase lá.")
	default:
		fmt.Println("Fim de semana!")
	}

	// switch sem variável (substituto de if/else aninhados)
	hora := time.Now().Hour()
	fmt.Print("\nSaudação pelo horário: ")
	switch {
	case hora < 12:
		fmt.Println("Bom dia!")
	case hora < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}
}
