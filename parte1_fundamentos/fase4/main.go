package main

import (
	"errors"
	"fmt"
)

// 1. Função Básica e Múltiplos Retornos
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero não permitida")
	}
	return a / b, nil // nil significa "sem erros" em Go (é o zero value para interfaces/ponteiros)
}

// 2. Retornos Nomeados (Named Return Values)
func calcularEstatisticas(notas []float64) (soma float64, media float64) {
	if len(notas) == 0 {
		return // "Naked return": vai retornar os zero values atuais (0.0, 0.0)
	}

	for _, nota := range notas {
		soma += nota
	}
	media = soma / float64(len(notas))

	return // Naked return implícito: retorna soma e media atuais
}

// 3. Closures
// Função que retorna outra função (um gerador)
func contador() func() int {
	i := 0
	return func() int { // Esta função anônima lembra o estado de 'i'
		i++
		return i
	}
}

func main() {
	fmt.Println("=== Fase 4: Funções e Comportamento ===\n")

	// 1. Múltiplos Retornos (e tratamento de erros base)
	fmt.Println("[Múltiplos Retornos]")
	resultado, err := dividir(10, 2)
	if err != nil { // Padrão Go idiomático
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Resultado com sucesso: %.2f\n", resultado)
	}

	// Simulação de Erro
	_, err2 := dividir(10, 0) // O _ ignora o valor numérico
	if err2 != nil {
		fmt.Println("Erro esperado:", err2)
	}

	// 2. Retornos Nomeados
	fmt.Println("\n[Retornos Nomeados]")
	minhasNotas := []float64{7.5, 8.0, 10.0}
	total, med := calcularEstatisticas(minhasNotas)
	fmt.Printf("Soma: %.2f | Média: %.2f\n", total, med)

	// 3. Funções anônimas e Closures
	fmt.Println("\n[Closures]")
	gerarProximoID := contador()
	fmt.Printf("ID: %d\n", gerarProximoID()) // 1
	fmt.Printf("ID: %d\n", gerarProximoID()) // 2
	fmt.Printf("ID: %d\n", gerarProximoID()) // 3

	// Criando um novo contador (estado é isolado)
	novoContador := contador()
	fmt.Printf("Novo Contador ID: %d\n", novoContador()) // 1

	// 4. A palavra-chave defer
	fmt.Println("\n[O uso do Defer]")
	executaSimulacaoDefer()
}

func executaSimulacaoDefer() {
	fmt.Println("Iniciando rotina de abertura...")
	
	// O defer programa esta função para rodar no final de executaSimulacaoDefer.
	// Eles são resolvidos em LIFO (Last In, First Out) caso haja mais de um.
	defer func() {
		fmt.Println("[DEFER] Limpando e fechando o que foi aberto.")
	}()

	fmt.Println("Realizando trabalho pesado e lógica de negócio.")
	fmt.Println("Processamento concluído. Retornando em breve...")
}
