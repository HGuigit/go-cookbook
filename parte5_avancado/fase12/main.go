package main

import "fmt"

// 1. Função Generic Básica
// A Constraint 'any' significa que T pode ser qualquer tipo.
func FiltrarPrimeiroElemento[T any](lista []T) T {
	if len(lista) == 0 {
		var zeroValue T // Cria e retorna o zero value do tipo genérico caso vazio
		return zeroValue
	}
	return lista[0]
}

// 2. Type Constraints Personalizadas
// Permite que N seja apenas um int ou float64.
type Number interface {
	int | float64
}

// O compilador agora sabe que N é numérico e permite a operação '+' e '/'
func Somar[N Number](a, b N) N {
	return a + b
}

// 3. Estruturas Genéricas (Structs)
// Útil para criar Árvores, Pilhas, Fila, e Respostas de API.
type Pilha[T any] struct {
	itens []T
}

func (p *Pilha[T]) Push(item T) {
	p.itens = append(p.itens, item)
}

func (p *Pilha[T]) Pop() T {
	if len(p.itens) == 0 {
		var zero T
		return zero
	}
	ultimoIndex := len(p.itens) - 1
	item := p.itens[ultimoIndex]
	p.itens = p.itens[:ultimoIndex]
	return item
}

func main() {
	fmt.Println("=== Fase 12: Generics ===\n")

	// Usando a Função Generic Básica
	fmt.Println("[Funções Genéricas e Any]")
	textos := []string{"Go", "Rust", "C++"}
	numeros := []int{42, 10, 5}

	// O compilador infere que T é string
	fmt.Println("Primeiro Texto:", FiltrarPrimeiroElemento(textos))
	
	// O compilador infere que T é int
	fmt.Println("Primeiro Número:", FiltrarPrimeiroElemento(numeros))

	// Usando Type Constraints
	fmt.Println("\n[Type Constraints]")
	fmt.Printf("Soma de int: %d\n", Somar(10, 20))
	fmt.Printf("Soma de float: %.2f\n", Somar(99.9, 0.1))
	// Somar("A", "B") // ERRO DE COMPILAÇÃO! String não atende à constraint Number.

	// Usando Estruturas Genéricas
	fmt.Println("\n[Estruturas de Dados Genéricas]")
	// Pilha que armazena APENAS strings
	pilhaNomes := Pilha[string]{}
	pilhaNomes.Push("Alice")
	pilhaNomes.Push("Bob")
	fmt.Println("Pop da Pilha (Strings):", pilhaNomes.Pop())

	// A mesmíssima estrutura para lidar com float64, de forma tipada (type-safe)
	pilhaPrecos := Pilha[float64]{}
	pilhaPrecos.Push(10.50)
	pilhaPrecos.Push(99.99)
	fmt.Println("Pop da Pilha (Floats):", pilhaPrecos.Pop())
}
