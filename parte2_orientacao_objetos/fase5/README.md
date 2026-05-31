# Fase 5: Structs, Métodos, Interfaces e Erros

Go não é uma linguagem Orientada a Objetos clássica (como Java ou C#). Não temos o conceito de `class` ou de "herança" (`extends`). Em vez disso, Go utiliza composição e interfaces implícitas para criar código altamente flexível e coeso.

## Structs (Estruturas)

Uma `struct` agrupa variáveis sob um único tipo. É o equivalente mais próximo que o Go possui de uma "Classe" para armazenamento de estado.

### Struct Tags
Muitas vezes, uma Struct precisa ser serializada ou desserializada (ex: convertida para JSON). O Go usa "Struct Tags" (strings colocadas logo após a declaração do tipo do campo) para instruir pacotes como o `encoding/json` ou `gorm` sobre como tratar aquela propriedade.

## Métodos

Em Go, você pode atrelar uma função a qualquer tipo que você tenha criado no seu pacote. Essa função passa a ser chamada de **Método**.
A sintaxe coloca um "Receiver" (recebedor) antes do nome do método.

### Ponteiros vs Valores no Receiver
- Se você quiser modificar os campos da `struct` original dentro do seu método, você **precisa** passar um ponteiro como receiver (`*MinhaStruct`).
- Se você não for modificar nada e a struct for pequena, passe por valor. Contudo, em estruturas grandes, passar por ponteiro evita cópias desnecessárias na memória.

## Interfaces Implícitas (Duck Typing)

Uma Interface define um conjunto de assinaturas de métodos. 
A grande sacada do Go é que **você não precisa declarar explicitamente que uma Struct implementa uma Interface** (não existe a palavra `implements`). Se uma struct possui os mesmos métodos com a exata assinatura exigida pela Interface, ela automaticamente a implementa. 
*Se anda como um pato, nada como um pato e grasna como um pato, então é um pato.*

## Tratamento de Erros e Empacotamento

Em Go, erros são apenas valores que implementam a interface built-in `error`.
O padrão é retornar o erro como o último valor da função. Você checa usando `if err != nil`. 

Para contextualizar de onde o erro veio, o Go permite o "Wrapping" (empacotamento) de erros com o pacote `fmt` e o verbo `%w`, criando uma cadeia de erros para facilitar o rastreamento, mantendo o erro original inspecionável com `errors.Is` ou `errors.As`.

---

### Executando o exemplo prático

Explore as structs e o polimorfismo do Go abrindo e rodando o arquivo `main.go`:
```bash
go run main.go
```

## 🚀 Exercício Prático

Agora é a sua vez! Entre na pasta `exercicio/` e abra o arquivo `main.go`.
Seu desafio é:
1. Criar uma Struct `Carro` com seus campos.
2. Criar métodos para essa struct, prestando atenção em quando usar o Receiver como Ponteiro (para conseguir alterar o estado interno da struct, como a velocidade).

Para testar, rode: `go run exercicio/main.go`
