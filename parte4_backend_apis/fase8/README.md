# Fase 8: Integrações, Background Workers e Filas

Em sistemas reais, nem todo processamento deve ser feito durante o tempo de resposta da requisição HTTP de um usuário. Tarefas pesadas, envios de e-mail e processamento de imagens devem ser jogados para execução em background para não travar a API.

## Padrão Worker Pool

O **Worker Pool** é o padrão arquitetural em Go para controlar concorrência em tarefas massivas.
Em vez de abrirmos uma Goroutine indiscriminadamente para cada tarefa (o que, dependendo do volume, poderia derrubar um banco de dados downstream pelo excesso de conexões concorrentes), nós criamos um número fixo de "Workers" (trabalhadores).

1. O fluxo principal envia tarefas (Tasks) para um **Channel (Canal)**, que age como uma Fila In-Memory.
2. Vários *Workers* ficam lendo simultaneamente deste mesmo canal.
3. Se houver 5 workers, apenas 5 tarefas serão processadas ao mesmo tempo, não importando se há 100 mil tarefas na fila.
4. Quando um worker finaliza sua tarefa, ele puxa a próxima imediatamente.

## Graceful Shutdown (Interceptação de Sinais)

Quando você precisa atualizar seu servidor ou o Kubernetes desliga o seu Pod, ele envia um sinal do Sistema Operacional (ex: `SIGTERM`). 
Se o seu código fechar abruptamente, clientes podem receber erros 502 (Bad Gateway) e tarefas no meio da execução serão perdidas.

O Go permite interceptar esses sinais usando `os/signal` atrelado a um Canal. Quando o sinal chega, o seu código inicia um desligamento progressivo (Graceful Shutdown):
- Para de aceitar novas requisições.
- Aguarda as requisições em andamento terminarem de processar.
- Fecha conexões com o Banco de Dados ou Filas.
- Somente então desliga a aplicação.

---

### Executando o exemplo prático

Rode o worker pool e teste cancelar com `Ctrl+C` enquanto estiver processando:
```bash
go run main.go
```

## 🚀 Exercício Prático

Agora é a sua vez! Entre na pasta `exercicio/` e abra o arquivo `main.go`.
Seu desafio é:
1. Construir uma versão miniatura de um Worker Pool, mas agora focado em "Envio de E-mails".
2. Você precisará instanciar um WaitGroup, iniciar 2 workers paralelos e usar um Channel com Buffer para mandar as mensagens!

Para testar, rode: `go run exercicio/main.go`
