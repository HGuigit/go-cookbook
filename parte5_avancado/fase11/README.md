# Fase 11: Qualidade de Código e Testes

O ecossistema do Go já possui ferramentas de teste e benchmarking de altíssima qualidade integradas nativamente, sem necessidade de baixar frameworks de terceiros (como Jest ou JUnit).

A regra básica é: se o seu código está no arquivo `calculadora.go`, seus testes devem ficar no arquivo `calculadora_test.go` na mesma pasta. Para rodar todos os testes de um projeto, basta executar `go test ./...`.

## O Pacote `testing` e os Table-Driven Tests

O padrão ouro (Gold Standard) para escrever testes em Go é o **Table-Driven Test**. 
Em vez de escrever uma função gigante para cada cenário de teste (`TestSomaValoresPositivos`, `TestSomaValoresNegativos`, etc), criamos uma única função de teste que itera sobre um slice de Structs (a "tabela"). Cada struct representa um Cenário com seus inputs e o resultado esperado.
Isso torna a adição de novos cenários de teste extremamente rápida.

## Mocks e Interfaces

Como visto na Fase 5, Go utiliza Interfaces Implícitas. Para testar uma regra de negócio que depende do Banco de Dados, você passa uma interface `Repositorio` para a sua função de negócio. No código real, você injeta o Repositório de Postgres. No teste, você injeta um `MockRepositorio` que implementa a mesma interface, mas retorna dados falsos estáticos para isolar a sua regra de negócio do mundo externo.

## Benchmarks (`-bench`)

Além de testes unitários (`*testing.T`), o Go permite criar funções que mensuram performance com `*testing.B`. O runtime do Go vai executar a sua função repetidas vezes para descobrir quanto tempo cada operação demora na média e quanta memória está sendo alocada no processo (utilizando a flag `-benchmem`).

---

### Executando o exemplo prático

Como os testes requerem arquivos separados, as funções de teste estão no arquivo `main_test.go`. 
Para executar os testes e o benchmark, rode no terminal dentro desta pasta:

```bash
# Rodar os testes normais com verbosidade
go test -v

# Rodar os benchmarks medindo também as alocações de memória
go test -bench=. -benchmem
```

## 🚀 Exercício Prático

Agora é a sua vez! Entre na pasta `exercicio/` e você verá dois arquivos (`main.go` e `exercicio_test.go`).
Seu desafio é:
1. Abrir o `exercicio_test.go` e implementar um **Table-Driven Test** para a função `ClassificaNota()`.
2. Rodar os testes para validar se você escreveu o cenário corretamente.

Para testar, rode: `cd exercicio` e depois `go test -v`
