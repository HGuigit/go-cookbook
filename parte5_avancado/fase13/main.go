package main

import (
	"bytes"
	"fmt"
	"sync"
	// Import anônimo ativaria a rota /debug/pprof/ se usarmos o DefaultServeMux
	// _ "net/http/pprof"
)

func main() {
	fmt.Println("=== Fase 13: Profiling e sync.Pool ===\n")

	// =========================================================
	// USO DO SYNC.POOL PARA ALTA PERFORMANCE
	// =========================================================
	
	// Imagina um cenário web onde cada Request gera um Buffer JSON de 10KB.
	// Se tivermos 10.000 requisições/segundo, estaríamos alocando e destruindo
	// 100MB de buffers a cada segundo. O Garbage Collector iria "fritar" a CPU.

	// A Solução: Criamos um Pool de Buffers.
	bufferPool := &sync.Pool{
		// A função New ensina o pool a criar um objeto NOVO caso ele esteja totalmente vazio
		New: func() any {
			fmt.Println("[Pool] -> Criou um NOVO *bytes.Buffer do zero na memória.")
			return new(bytes.Buffer)
		},
	}

	fmt.Println("--- Requisição do Usuário A ---")
	// Pegamos um buffer do Pool
	bufA := bufferPool.Get().(*bytes.Buffer)
	
	// Usamos o buffer...
	bufA.WriteString("Processando payload gigante do Usuário A...")
	fmt.Println("Buffer A:", bufA.String())
	
	// ANTES de devolver para o pool, DEVEMOS LIMPÁ-LO (Reset),
	// caso contrário, o próximo usuário leria os dados antigos.
	bufA.Reset()
	
	// Devolvemos ao pool para reciclagem (evitamos chamar o Garbage Collector)
	bufferPool.Put(bufA)

	fmt.Println("\n--- Requisição do Usuário B ---")
	// Como o bufA foi devolvido limpo, o Get vai REAPROVEITAR aquela memória exata.
	// O log da função "New" não será impresso!
	bufB := bufferPool.Get().(*bytes.Buffer)
	bufB.WriteString("Payload do Usuário B.")
	fmt.Println("Buffer B (Reaproveitado):", bufB.String())


	// =========================================================
	// PPROF - PROFILING DE APLICAÇÃO
	// =========================================================
	// Para testar o pprof na vida real, descomente o bloco abaixo.
	// O código ficará rodando. Abra outro terminal e digite:
	// go tool pprof http://localhost:6060/debug/pprof/heap
	
	/*
	fmt.Println("\n[!] Iniciando servidor pprof na porta 6060...")
	fmt.Println("Acesse http://localhost:6060/debug/pprof/")
	// Bloqueia e serve a porta de debug
	log.Fatal(http.ListenAndServe("localhost:6060", nil))
	*/
}
