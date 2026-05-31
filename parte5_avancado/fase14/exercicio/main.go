package main

import "fmt"

// TODO 1: Crie a Camada de Repositório (Interface "ProdutoRepository" com método Salvar)

// TODO 2: Crie a Camada de Implementação Falsa ("MockProdutoRepo" que implementa a interface e só imprime um log)

// TODO 3: Crie a Camada de Serviço ("ProdutoService" que recebe a Interface no construtor)
// Adicione a regra de negócio: Se o preço do produto for menor que 0, retorne um erro, senão chame repo.Salvar()

func main() {
	fmt.Println("--- Exercício: Clean Architecture (Injeção) ---")

	// TODO 4: Instancie o MockProdutoRepo e passe ele no construtor do ProdutoService!
	
	// TODO 5: Simule o cadastro de um produto com preço negativo e veja se o erro volta corretamente sem chamar o banco.
}
