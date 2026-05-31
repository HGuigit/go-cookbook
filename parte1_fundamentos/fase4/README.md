# Fase 4: Funções e Comportamento

As funções em Go são "cidadãs de primeira classe", o que significa que podem ser passadas como argumentos, atribuídas a variáveis e retornadas por outras funções.

## Múltiplos Retornos

Esta é uma das marcas registradas do Go e é a base de todo o modelo de **Tratamento de Erro** da linguagem. Uma função pode retornar mais de um valor simultaneamente, o que dispensa a criação de estruturas ou classes complexas apenas para agrupar resultados (como seria feito em Java com um objeto contendo "Resultado" e "MensagemErro").

## Retornos Nomeados (Named Return Values)

Você pode dar nomes aos retornos na própria assinatura da função. Ao fazer isso, o Go inicializa as variáveis com seus Zero Values correspondentes e permite que você use a instrução `return` pura (o "naked return"), que vai retornar os valores atuais daquelas variáveis. 

> [!TIP]
> Embora seja útil para clarear a documentação (a assinatura por si só já diz o que cada variável retornada significa), os "naked returns" em funções longas podem prejudicar a legibilidade, pois você terá que procurar no corpo da função onde os valores finais foram definidos.

## Funções Anônimas e Closures

Você pode declarar funções que não possuem nome diretamente dentro do escopo de outras funções. Essas funções têm acesso às variáveis do escopo que as envolveu (isto é o que chamamos de **Closure**).

## O Poder do `defer`

A instrução `defer` empilha a execução de uma função para o **exato momento antes** da função atual retornar. 

É o mecanismo canônico e indispensável no Go para garantir limpezas de recursos, tais como:
- Fechar conexões com Banco de Dados.
- Fechar arquivos recém-abertos (arquivos ou *file descriptors*).
- Desbloquear Mutexes (Concorrência).
- Fazer *recover* de um panic (evitando que a aplicação morra abruptamente).

Ao utilizar o `defer`, você programa o "fechamento" logo abaixo da "abertura", deixando seu código seguro e impossível de esquecer o fechamento, mesmo que a função possua múltiplos `return` no meio.

---

### Executando o exemplo prático

Rode os conceitos de funções abrindo e executando o arquivo da fase 4:
```bash
go run main.go
```

## 🚀 Exercício Prático

Agora é a sua vez! Entre na pasta `exercicio/` e abra o arquivo `main.go`.
Seu desafio é:
1. Escrever uma função que possui múltiplos retornos (Área e Perímetro de um retângulo).
2. Utilizar o `defer` para garantir a execução de um print ao final de uma simulação de servidor.

Para testar, rode: `go run exercicio/main.go`
