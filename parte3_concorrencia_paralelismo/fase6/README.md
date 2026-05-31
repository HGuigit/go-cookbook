# Fase 6: Concorrência - O Poder do Go

A maior vantagem do Go em relação a outras linguagens é o seu modelo de concorrência. Ele foi desenhado para utilizar todo o potencial de processadores multi-core de forma simples e natural, através das Goroutines e Channels.

## Goroutines

Goroutines são funções executadas concorrentemente com outras funções. Elas não são "threads" de sistema operacional (OS Threads). São threads gerenciadas pelo próprio *runtime* do Go (conhecidas como "Green Threads"). 
Você pode abrir milhares ou até milhões de Goroutines simultaneamente no mesmo programa com uso baixíssimo de memória (elas começam com apenas ~2KB).

A sintaxe é extremamente simples: basta colocar a palavra `go` antes da chamada da função.

## Channels (Canais)

"Não se comunique compartilhando memória; em vez disso, compartilhe memória se comunicando." Esse é o lema do Go. 
Channels são os tubos que conectam as goroutines. Você pode enviar valores para canais de uma goroutine e recebê-los em outra, sincronizando a execução naturalmente de forma Thread-Safe.

## O Padrão `select`

Quando sua Goroutine precisa lidar com **vários canais** simultaneamente, o `select` entra em cena. Funciona como um `switch`, mas para operações de leitura/escrita em canais. Ele espera até que um dos canais esteja pronto para processar.

## O Pacote `context`

No desenvolvimento de servidores web e microserviços, é muito comum precisarmos cancelar tarefas concorrentes que estão demorando demais ou quando a conexão do usuário foi perdida. O pacote `context` é usado em todo o ecossistema Go para gerenciar tempo limite (Timeouts) e enviar sinais de cancelamento pelas camadas do software.

## Sincronização Explícita (`sync`)

Embora canais sejam o padrão ouro, às vezes precisamos de formas tradicionais de proteger um recurso de leituras simultâneas ou apenas aguardar rotinas terminarem:
- **`sync.WaitGroup`**: Espera que um conjunto de Goroutines termine a sua execução antes de liberar o fluxo principal.
- **`sync.Mutex`**: Tranca uma variável na memória (Lock/Unlock) garantindo que apenas uma goroutine acesse o recurso por vez (prevenindo Data Races).

---

### Executando o exemplo prático

Para ver a verdadeira força da concorrência, veja o código da fase 6:
```bash
go run main.go
```

## 🚀 Exercício Prático

Agora é a sua vez! Entre na pasta `exercicio/` e abra o arquivo `main.go`.
Seu desafio é:
1. Usar as famosas `goroutines` para executar uma tarefa em paralelo ao fluxo principal.
2. Fazer com que a `goroutine` envie um valor de volta para a `main` através de um `channel`, para que a thread principal só encerre quando a goroutine terminar.

Para testar, rode: `go run exercicio/main.go`
