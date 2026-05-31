package main

import "fmt"

func main() {
	fmt.Println("=== Fase 3: Coleções e Estruturas de Dados ===\n")

	// 1. Arrays (Tamanho Fixo)
	fmt.Println("[Arrays]")
	var notas [3]float64
	notas[0] = 8.5
	notas[1] = 9.0
	notas[2] = 7.5
	fmt.Printf("Array de notas: %v | Tamanho: %d\n", notas, len(notas))

	// 2. Slices (Dinâmicos)
	fmt.Println("\n[Slices]")
	// Sintaxe de Slice Literal (Parece um array sem tamanho definido)
	linguagens := []string{"Go", "Python", "Rust"}
	fmt.Printf("Slices: %v | Len: %d | Cap: %d\n", linguagens, len(linguagens), cap(linguagens))

	// Como o append funciona e mexe com a Capacidade
	fmt.Println("\n[Append e Capacidade]")
	var numeros []int // Slice com len=0 e cap=0

	for i := 1; i <= 5; i++ {
		numeros = append(numeros, i)
		fmt.Printf("Adicionado %d -> Len: %d | Cap: %d\n", i, len(numeros), cap(numeros))
	}
	// Note no output como a capacidade dobra (0, 1, 2, 4, 8) para evitar realocações a cada insert.

	// Inicialização otimizada com make
	// make([]tipo, length, capacity)
	otimizado := make([]int, 0, 100)
	fmt.Printf("\nSlice otimizado -> Len: %d | Cap: %d\n", len(otimizado), cap(otimizado))

	// 3. Maps (Chave / Valor)
	fmt.Println("\n[Maps]")
	// Inicializando com make
	idades := make(map[string]int)
	idades["Alice"] = 25
	idades["Bob"] = 30

	// Inicialização com Map Literal
	paises := map[string]string{
		"BR": "Brasil",
		"US": "Estados Unidos",
	}
	fmt.Printf("Idade da Alice: %d\n", idades["Alice"])
	fmt.Printf("País BR: %s\n", paises["BR"])

	// Checando existência da chave no Map (o padrão "comma ok")
	fmt.Println("\n[Verificando chaves no Map]")
	
	valor, ok := idades["Carlos"] // Carlos não existe
	if ok {
		fmt.Printf("Idade do Carlos é %d\n", valor)
	} else {
		fmt.Println("A chave 'Carlos' não existe no map. Valor retornado (zero value):", valor)
	}

	// Removendo itens de um Map
	delete(paises, "US")
	fmt.Println("Mapa de países após exclusão de 'US':", paises)
}
