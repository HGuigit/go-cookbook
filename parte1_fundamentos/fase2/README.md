# Fase 2: Estruturas de Controle de Fluxo

A filosofia de design do Go preza pelo minimalismo e legibilidade. Por isso, a linguagem possui um número reduzido de estruturas de controle, que são extremamente versáteis e cobrem todas as necessidades lógicas que você teria com laços diversos em outras linguagens (como `while`, `do-while` ou `foreach`).

## `if` e `else`

A sintaxe de condicional básica do Go dispensa o uso de parênteses ao redor da condição. Além disso, as chaves `{}` são obrigatórias, mesmo para blocos com uma única instrução.

### Short Statement no `if`
Um recurso poderoso do Go é a capacidade de declarar uma variável dentro da própria instrução `if` (conhecido como Short Statement). Essa variável terá seu escopo limitado aos blocos do `if` e `else`, não poluindo o escopo externo.

## O Único Laço de Repetição: `for`

O Go removeu completamente comandos como `while` e `do-while`. A única estrutura de repetição disponível é o `for`. A versatilidade dele permite três abordagens distintas:

1. **`for` Tradicional**: Similar à linguagem C (`for inicializacao; condicao; pos-execucao`).
2. **`for` como `while`**: Você omite a inicialização e a pós-execução, passando apenas a condição.
3. **`for` infinito**: Omissão de todas as expressões. Utilizado frequentemente junto com canais e goroutines (assuntos de concorrência mais avançados).

### Iterações com `range`

A palavra-chave `range` é utilizada conjuntamente com o `for` para iterar sobre estruturas de dados iteráveis, como slices, arrays, maps, strings ou canais. Ele retorna dois valores na maioria dos casos (índice e valor, ou chave e valor para maps).

> [!TIP]
> Se você precisar apenas do valor (ignorando o índice) no `range`, você deve omitir a primeira variável utilizando o `Blank Identifier` (o underscore `_`), já que o Go proíbe variáveis declaradas não utilizadas.

## O Poderoso `switch`

O `switch` do Go foi modernizado:
- **Não há fallthrough automático**: Em linguagens como C ou Java, se você omitir um `break`, a execução continua ("cai") para o próximo `case`. No Go, isso **não acontece**. O `switch` para assim que um `case` for avaliado e seu bloco executado. Para forçar esse comportamento, existe a instrução `fallthrough`.
- **Casos com expressões booleanas**: Você pode criar um `switch` sem passar uma variável, avaliando condições dentro de cada `case`, servindo como uma alternativa mais limpa e legível para correntes longas de `if-else if`.

---

### Executando o exemplo prático

Abra o arquivo `main.go` desta pasta e experimente as diferentes formas de controle de fluxo rodando:
```bash
go run main.go
```

## 🚀 Exercício Prático

Agora é a sua vez! Entre na pasta `exercicio/` e abra o arquivo `main.go`.
Seu desafio é:
1. Criar um laço `for` que itere de 1 a 10.
2. Dentro do laço, usar um `if` ou `switch` para checar e imprimir se o número é Par ou Ímpar.

Para testar, rode: `go run exercicio/main.go`
