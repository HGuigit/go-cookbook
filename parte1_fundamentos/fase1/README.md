# Fase 1: Introdução, Variáveis e Tipos Básicos

O objetivo desta fase é entender a sintaxe fundamental da linguagem Go, sua tipagem estática e como o gerenciamento inicial de memória e inicialização funciona.

## Declaração de Variáveis

Go oferece duas formas principais de declarar variáveis:

1.  **Declaração curta (`:=`)**: Utilizada apenas **dentro de funções**. O Go infere o tipo da variável automaticamente.
2.  **Palavra-chave `var`**: Utilizada para declaração explícita. Necessária quando queremos declarar variáveis em escopo global (nível de pacote) ou quando não queremos inicializar a variável imediatamente.

## Tipos Primitivos

Os tipos de dados mais fundamentais em Go incluem:
- `string`: Sequência de caracteres (texto).
- `int`, `int32`, `int64`: Números inteiros. Normalmente, usa-se apenas `int` (cujo tamanho depende da arquitetura do SO, 32 ou 64 bits).
- `float32`, `float64`: Números de ponto flutuante. O padrão da declaração curta infere sempre `float64`.
- `bool`: Verdadeiro (`true`) ou Falso (`false`).

## O "Zero Value"

Em Go, **não existem variáveis não inicializadas** contendo "lixo" de memória ou `undefined`. Se você declarar uma variável e não atribuir um valor, ela recebe o valor zero (Zero Value) de seu respectivo tipo:
- `int` / `float64`: `0`
- `string`: `""` (string vazia)
- `bool`: `false`
- Ponteiros, funções, interfaces, slices, canais e maps: `nil`

> [!NOTE]
> O Zero Value reduz consideravelmente bugs em produção causados por uso de variáveis não inicializadas.

## Constantes e `iota`

Constantes (`const`) são valores imutáveis. O construtor `iota` é um identificador especial utilizado para criar **enums** (sequências numéricas autoincrementais) em blocos de constantes de maneira simplificada.

---

### Executando o exemplo prático

Veja o código em `main.go` nesta pasta para explorar todos estes conceitos na prática. Para rodar:
```bash
go run main.go
```

## 🚀 Exercício Prático

Agora é a sua vez! Entre na pasta `exercicio/` e abra o arquivo `main.go`.
Seu desafio é:
1. Declarar variáveis (usando `:=` e `var`) para armazenar dados de um 'Produto' (ex: nome, preço e se está em estoque).
2. Imprimir os valores formatados no terminal.

Para testar, rode: `go run exercicio/main.go`
