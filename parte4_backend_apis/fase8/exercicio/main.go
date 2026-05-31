package main

import (
	"fmt"
	"sync"
)

// TODO 1: Crie a função do "worker" (que simula o envio de email)
// Assinatura sugerida: func enviarEmail(id int, filaEmails <-chan string, wg *sync.WaitGroup)
// Dentro da função, itere (usando range) lendo os emails da fila e imprima algo como "[Worker X] Enviando email para Y"

func main() {
	fmt.Println("--- Exercício: Workers e Filas ---")

	// TODO 2: Crie um canal com buffer de 5 posições chamado 'filaEmails'

	// TODO 3: Instancie um sync.WaitGroup

	// TODO 4: Suba 2 workers em Goroutines concorrentes, passando o canal e o WaitGroup

	// TODO 5: Envie 5 endereços de email (strings) para o canal 'filaEmails'

	// TODO 6: Feche a fila (close) para que os workers saibam que acabou, e dê wg.Wait() para aguardar
}
