# Fase 3: Coleções e Estruturas de Dados

Para armazenar e agrupar dados, o Go fornece três estruturas principais: Arrays, Slices e Maps.

## Arrays

Em Go, **Arrays têm tamanho fixo** e seu tamanho faz parte do seu tipo. Isso significa que um `[3]int` é um tipo completamente diferente de um `[4]int`. Por causa da limitação de tamanho estático, Arrays raramente são usados diretamente na programação cotidiana em Go.

## Slices (Fatias)

Slices são a estrutura de dados mais importante e utilizada no Go. Eles são "janelas" dinâmicas que apontam para um Array oculto (underlying array) por debaixo dos panos.

Diferente dos Arrays, Slices têm um tamanho dinâmico e crescem sob demanda.

### Como funciona o `append`?

Slices possuem três propriedades internas:
- **Pointer**: Aponta para o array underlying.
- **Length (Tamanho)**: O número de elementos atualmente no slice. Acessível via `len(s)`.
- **Capacity (Capacidade)**: O número máximo de elementos que o array underlying suporta antes de precisar ser redimensionado na memória. Acessível via `cap(s)`.

Quando você usa a função built-in `append(slice, elemento)`, o Go verifica se ainda há Capacidade. Se sim, ele apenas insere o valor e incrementa o Tamanho. Se a Capacidade esgotou, o Go **aloca um novo array (geralmente com o dobro da capacidade)**, copia os dados antigos, insere o novo elemento e atualiza o ponteiro do Slice.

> [!WARNING]
> Entender Capacidade vs Tamanho é fundamental para evitar bugs de memória e perda de performance causadas por muitas realocações no `append`. Se você sabe de antemão o tamanho aproximado da coleção, inicie o slice com a capacidade correta usando `make([]tipo, tamanho, capacidade)`.

## Maps (Dicionários / Hash Tables)

Maps são a estrutura para chave/valor do Go, similares a dicionários em Python ou Objetos no Javascript (sem métodos). Eles são não-ordenados. A declaração padrão é `map[TipoChave]TipoValor`.

### Checando a existência de uma chave
Acessar uma chave inexistente em um map retorna o Zero Value do tipo do valor (ex: `0` para inteiros). Para ter certeza se a chave existe ou se o valor apenas é de fato zero, usamos o retorno duplo (conhecido como padrão "comma ok").

---

### Executando o exemplo prático

Abra o arquivo `main.go` desta pasta para testar slices e maps rodando:
```bash
go run main.go
```

## 🚀 Exercício Prático

Agora é a sua vez! Entre na pasta `exercicio/` e abra o arquivo `main.go`.
Seu desafio é:
1. Criar um Slice de frutas e adicionar mais itens usando `append`.
2. Criar um Map com preços e usar `for range` para imprimir todos eles.

Para testar, rode: `go run exercicio/main.go`
