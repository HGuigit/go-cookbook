package main

import "fmt"

// Constantes globais
const AppName = "Go Cookbook"
const Version = 1.0

// Usando iota para criar um "Enum" simplificado
const (
	Domingo = iota // 0
	Segunda        // 1
	Terca          // 2
	Quarta         // 3
	Quinta         // 4
	Sexta          // 5
	Sabado         // 6
)

func main() {
	fmt.Println("---", AppName, "v", Version, "---")

	// 1. O "Zero Value"
	// Declaração explícita com var sem valor inicial
	var semValorInt int
	var semValorString string
	var semValorBool bool

	fmt.Println("\n[Zero Values]")
	fmt.Printf("Int: %d | String: %q | Bool: %t\n", semValorInt, semValorString, semValorBool)

	// 2. Declaração curta de variáveis (:=)
	// Somente permitido dentro de funções. Infere o tipo.
	nome := "Desenvolvedor Go"
	idade := 30
	altura := 1.75 // infere como float64

	fmt.Println("\n[Declaração Curta]")
	fmt.Printf("Nome: %s (Tipo: %T)\n", nome, nome)
	fmt.Printf("Idade: %d (Tipo: %T)\n", idade, idade)
	fmt.Printf("Altura: %.2f (Tipo: %T)\n", altura, altura)

	// 3. Trabalhando com tipos primitivos explicitamente
	var numeroPequeno int32 = 100
	var preco float32 = 19.99

	fmt.Println("\n[Tipos Explícitos]")
	fmt.Printf("Número: %d (Tipo: %T)\n", numeroPequeno, numeroPequeno)
	fmt.Printf("Preço: %.2f (Tipo: %T)\n", preco, preco)

	// 4. Demonstrando o Enum com iota
	fmt.Println("\n[Constantes e iota (Enum)]")
	fmt.Printf("Domingo: %d\n", Domingo)
	fmt.Printf("Quarta: %d\n", Quarta)
	fmt.Printf("Sábado: %d\n", Sabado)
}
