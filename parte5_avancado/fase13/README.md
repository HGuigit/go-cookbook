# Fase 13: Profiling e Análise Profunda (pprof)

Quando sua aplicação está em produção e você nota um aumento bizarro no consumo de CPU ou memória (Memory Leaks), o Go oferece uma das ferramentas mais incríveis nativamente: o pacote `net/http/pprof`.

O `pprof` permite que você acesse endpoints HTTP ocultos no seu servidor para baixar "retalhos" do uso atual da CPU, memória e bloqueios de mutexes de um sistema **rodando em produção**, com um overhead baixíssimo.

## Habilitando o pprof
Se você usa o `http.DefaultServeMux`, basta fazer um import anônimo do pacote:
`import _ "net/http/pprof"`

Isso habilitará rotas mágicas sob `/debug/pprof/`.

## Analisando Dumps
Após baixar o arquivo de dump (ex: `go tool pprof http://localhost:8080/debug/pprof/heap`), você ganha um terminal interativo onde pode usar comandos como:
- `top`: Mostra as funções que estão gastando mais memória no exato momento.
- `web`: (Requer Graphviz) Gera um SVG visual lindíssimo mostrando o fluxo de chamadas e quem é o ofensor do vazamento de memória.

## Reaproveitamento de Memória (`sync.Pool`)
Quando o pprof mostra que o Gargalo do seu sistema é o Garbage Collector rodando constantemente (porque você está alocando e jogando fora muitas variáveis locais gigantes a cada request HTTP), nós usamos o padrão de projeto chamado **Object Pool**.

O `sync.Pool` mantém objetos que não estão mais em uso (mas já foram alocados na memória Heap) para que a próxima Goroutine possa "reciclar" a variável, em vez de exigir que o Go aloque nova memória no S.O. e acione o Garbage Collector depois.

---

### Executando o exemplo prático

Rode o exemplo do `sync.Pool` em `main.go`. Se você quiser expor a API do pprof e brincar no terminal, descomente o bloco indicado no código.
```bash
go run main.go
```

## 🚀 Exercício Prático

Agora é a sua vez! Entre na pasta `exercicio/` e abra o arquivo `main.go`.
Seu desafio é:
1. Criar um `sync.Pool` configurado para retornar structs do tipo `SessaoUsuario`.
2. Pegar uma struct do pool, preencher, e depois fazer o "Reset" (limpando o token) antes de devolvê-la para a piscina de memória!

Para testar, rode: `go run exercicio/main.go`
