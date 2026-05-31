package main

import "fmt"

func ClassificaNota(nota float64) string {
	if nota >= 7.0 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println("--- Exercício: Testes (main.go) ---")
	fmt.Println("Abra o arquivo exercicio_test.go para completar o desafio!")
}
