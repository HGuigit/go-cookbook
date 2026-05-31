package main

import (
	"fmt"
	"sync"
)

type SessaoUsuario struct {
	ID    int
	Token string
}

func main() {
	fmt.Println("--- Exercício: sync.Pool ---")

	// TODO 1: Crie um sync.Pool chamado "poolSessoes"
	// Na função "New" do pool, você deve retornar um ponteiro novo: &SessaoUsuario{}

	// TODO 2: Pegue uma Sessao do Pool usando poolSessoes.Get().(*SessaoUsuario)
	
	// TODO 3: Preencha com ID=1 e Token="ABC" e imprima no terminal.

	// TODO 4: Zere os campos da struct (Reset manual) para não vazar dados!
	
	// TODO 5: Devolva para o pool usando poolSessoes.Put(...)
}
