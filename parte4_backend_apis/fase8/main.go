package main

import (
	"context"
	"fmt"
	"math/rand"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// 1. Worker Pool
// O worker fica em loop infinito consumindo tarefas do canal até ele ser fechado.
func worker(id int, filaTarefas <-chan string, filaResultados chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	
	// Usamos range no canal para processar enquanto ele não for fechado (close())
	for tarefa := range filaTarefas {
		fmt.Printf("[Worker %d] Iniciou: %s\n", id, tarefa)
		
		// Simula tempo de processamento da tarefa
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		
		filaResultados <- fmt.Sprintf("Concluído por W%d: %s", id, tarefa)
	}
	fmt.Printf("[Worker %d] Encerrado pois a fila de tarefas fechou.\n", id)
}

func main() {
	fmt.Println("=== Fase 8: Background Workers e Graceful Shutdown ===\n")

	// Canais com Buffer (para não travar o for do fluxo principal ao preencher as filas)
	filaTarefas := make(chan string, 10)
	filaResultados := make(chan string, 10)

	var wg sync.WaitGroup

	// Iniciando um pool fixo de 3 workers
	fmt.Println("-> Iniciando Pool de 3 Workers...")
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, filaTarefas, filaResultados, &wg)
	}

	// Rotina para escutar sinais de interrupção (Graceful Shutdown) Moderno (Go 1.16+)
	// O Context será cancelado automaticamente se receber SIGINT (Ctrl+C) ou SIGTERM (Docker/K8s)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Enviando 5 tarefas para a fila concorrentemente com o resto
	go func() {
		for i := 1; i <= 5; i++ {
			filaTarefas <- fmt.Sprintf("Email_Lote_%d", i)
		}
		// Fechando o canal após enviar todas as tarefas. 
		fmt.Println("-> Todas as tarefas enfileiradas. Fechando Fila de Tarefas.")
		close(filaTarefas)
	}()

	// Fechando canal de resultados SOMENTE após todos os workers terminarem
	go func() {
		wg.Wait()
		close(filaResultados)
	}()

	// Loop principal consumindo os resultados e escutando por cancelamentos
	fmt.Println("-> Consumindo resultados ou aguardando sinal do Sistema...\n")
	ativo := true
	for ativo {
		select {
		case res, ok := <-filaResultados:
			if !ok {
				// Canal de resultados fechou. Tudo processou.
				fmt.Println("\n-> Processamento total finalizado com sucesso!")
				ativo = false
			} else {
				fmt.Println(">>> RESULTADO:", res)
			}

		case <-ctx.Done():
			// Se o usuário der Ctrl+C ou o kubernetes cancelar
			fmt.Println("\n[!] SINAL DE DESLIGAMENTO RECEBIDO (Graceful Shutdown iniciado)...")
			fmt.Println("Na vida real: pararíamos de aceitar requests aqui e esperaríamos as rotinas terminarem.")
			ativo = false // Rompe o for
		}
	}
	
	fmt.Println("Desligando aplicação de forma segura.")
}
