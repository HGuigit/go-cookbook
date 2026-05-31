package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Processamento fake para simular uso de CPU/Rede
func fazerTrabalhoPesado(id int, ch chan<- string) {
	time.Sleep(time.Duration(id) * 200 * time.Millisecond)
	ch <- fmt.Sprintf("Goroutine %d finalizou o trabalho!", id) // Escreve no canal
}

func main() {
	fmt.Println("=== Fase 6: Concorrência e Paralelismo ===\n")

	// 1. Goroutines e Channels Básicos
	fmt.Println("[Goroutines e Canais]")
	canalNotificacoes := make(chan string) // Canal de strings

	go fazerTrabalhoPesado(1, canalNotificacoes)
	go fazerTrabalhoPesado(2, canalNotificacoes)

	// Como disparamos duas rotinas, vamos aguardar as duas mensagens no canal.
	// A leitura de um canal BLOQUEIA a execução até a mensagem chegar.
	msg1 := <-canalNotificacoes // Lê do canal
	msg2 := <-canalNotificacoes
	fmt.Println(msg1)
	fmt.Println(msg2)

	// 2. Sincronização com sync.WaitGroup
	fmt.Println("\n[sync.WaitGroup (Aguardando todos terminarem)]")
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Adiciona 1 contador por Goroutine
		go func(id int) {
			defer wg.Done() // Decrementa 1 ao final
			fmt.Printf("Worker %d trabalhando...\n", id)
			time.Sleep(100 * time.Millisecond)
		}(i)
	}
	wg.Wait() // Trava o fluxo principal até o contador do wg chegar a zero
	fmt.Println("Todos os workers finalizaram!")

	// 3. O uso do Select (Multiplexação de canais)
	fmt.Println("\n[Select e Timeout]")
	canalRapido := make(chan string)
	
	go func() {
		time.Sleep(50 * time.Millisecond)
		canalRapido <- "Resposta Rápida API A"
	}()

	// Select aguarda no primeiro canal que cuspir dado
	select {
	case res := <-canalRapido:
		fmt.Println("Recebido com sucesso:", res)
	case <-time.After(30 * time.Millisecond): // Timeout simulado
		fmt.Println("Erro: A operação excedeu o limite de tempo!")
	}

	// 4. Uso de Context para Cancelamento Controlado
	fmt.Println("\n[Gerenciamento de Contexto]")
	
	// Cria um contexto que vai cancelar automaticamente em 100ms
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel() // Boa prática: libera o context quando a função pai terminar

	canalTrabalhoContexto := make(chan string)
	
	// Rotina longa de 300ms (O contexto vai cancelar antes disso)
	go func(c context.Context, ch chan string) {
		select {
		case <-time.After(300 * time.Millisecond):
			ch <- "Trabalho finalizado (demorou muito)"
		case <-c.Done():
			// c.Done() recebe um sinal quando o contexto expira ou é cancelado manualmente
			ch <- "Aviso: Trabalho interrompido! O contexto expirou: " + c.Err().Error()
		}
	}(ctx, canalTrabalhoContexto)

	fmt.Println(<-canalTrabalhoContexto)
}
