# Fase 12: Generics

Por muitos anos, o Go foi criticado pela falta de Genéricos. Isso forçava os desenvolvedores a criar a mesma função repetidas vezes para tipos diferentes (uma para somar arrays de `int`, outra para `float64`, outra para `int32`), ou apelar para o uso de `interface{}` vazio, que perdia a garantia de segurança de tipos em tempo de compilação.

A partir do **Go 1.18**, os Generics (Parâmetros de Tipo) foram introduzidos, permitindo a escrita de código altamente reutilizável, mantendo a Tipagem Forte estática.

## Sintaxe de Generics

O parâmetro de tipo é definido entre colchetes `[...]` antes dos argumentos regulares da função.
A declaração `func ImprimirArray[T any](arr []T)` informa ao compilador: "Existe um tipo `T` que aceita `any` (qualquer coisa). O argumento desta função é um slice desse tipo `T`".

Quando a função é chamada, o compilador do Go normalmente infere automaticamente o tipo com base no argumento passado, ou você pode passar explicitamente: `ImprimirArray[int](meuArray)`.

## Type Constraints (Restrições de Tipo)

A palavra-chave `any` é um alias para a interface vazia e permite *absolutamente qualquer tipo*. Contudo, você muitas vezes precisará restringir as opções.

Exemplo: se sua função Generic faz uma soma (`a + b`), ela não pode receber `any`, pois você não pode usar o operador matemático `+` em um `bool` ou `struct`. 
Para resolver isso, você cria **Constraints**. Uma Constraint é basicamente uma Interface, mas em vez de definir assinaturas de métodos, ela define uma **união de tipos primitivos permitidos** (ex: `int | float64 | float32`).

O pacote `golang.org/x/exp/constraints` da comunidade possui as restrições mais comuns prontas para uso (como `constraints.Ordered`, que engloba tudo que permite usar `>`, `<`, etc). O tipo embutido `comparable` engloba qualquer coisa que suporte `==` e `!=` (essencial para chaves de Maps).

---

### Executando o exemplo prático

Explore a redução drástica de repetição de código proporcionada pelos Generics abrindo o exemplo:
```bash
go run main.go
```

## 🚀 Exercício Prático

Agora é a sua vez! Entre na pasta `exercicio/` e abra o arquivo `main.go`.
Seu desafio é:
1. Declarar uma Constraint que suporte strings ou inteiros.
2. Escrever uma função genérica `SaoIguais` que verifique se duas variáveis do tipo `T` são idênticas usando `==`.

Para testar, rode: `go run exercicio/main.go`
