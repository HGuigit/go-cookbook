# Fase 14: Estrutura de Projetos (Clean Architecture)

Ao contrário de frameworks dogmáticos (como Rails ou Django), a linguagem Go não impõe uma estrutura de pastas estrita. Contudo, a comunidade adotou informalmente o **Standard Go Project Layout** para projetos grandes.

## Standard Project Layout

A estrutura mais comum de pastas em um projeto profissional Go:

- `cmd/`: Ponto de entrada da aplicação. Cada pasta dentro de `cmd/` deve ter um arquivo `main.go`. Ex: `cmd/api/main.go` (a API web), `cmd/worker/main.go` (o processador de filas).
- `internal/`: Código que não pode e não deve ser importado por outros repositórios no GitHub. O compilador do Go **impede** fisicamente que pacotes fora da raiz do projeto importem o que está na pasta `internal/`. É onde fica 95% do seu código.
- `pkg/`: Código utilitário que é seguro para ser exportado e usado por outras aplicações da sua empresa (ex: um parser customizado de CNPJ).

## Clean Architecture (Arquitetura Limpa / Hexagonal)

A premissa da Clean Architecture é proteger as regras de negócio de mudanças externas (Banco de Dados, Frameworks Web, Filas SQS, etc).
Isso é feito usando **Camadas** e **Injeção de Dependências via Interfaces**.

Em um projeto Go, dividimos o `internal/` em 3 camadas principais:

1. **Entity (Modelos Core)**: As structs puras (`Usuario{}`). Não sabem nada de JSON ou de SQL.
2. **Repository (Adapters - Saída)**: Implementa as queries no banco de dados (`UsuarioRepositoryImpl`).
3. **Service / Usecase**: A regra de negócio pura. **Este é o segredo**: O Service não recebe o banco de dados concreto, ele exige uma **Interface**. Assim, o negócio não sabe se os dados vêm do Postgres, MongoDB ou de um Mock no teste unitário.
4. **Handler / Controller (Adapters - Entrada)**: Ouve o HTTP/gRPC, faz parse de JSON e chama os métodos do Service.

Isso previne o Monolito Espaguete e permite escalar a aplicação mantendo a sanidade do código.

---

### Executando o exemplo prático

O arquivo da Fase 14 não possui um código executável trivial, mas sim **uma simulação arquitetural** de como montar os pacotes (handlers, services e repositórios) e conectá-los através de injeção de dependência na `main()`.

Abra `main.go` e estude o esqueleto de inicialização!

## 🚀 Exercício Prático

Agora é a sua vez! Entre na pasta `exercicio/` e abra o arquivo `main.go`.
Seu desafio é:
1. Montar a hierarquia de `Repository` -> `Service` para a entidade `Produto`.
2. Fazer a Injeção de Dependência usando a Interface, e depois simular uma regra de negócio na camada de Serviço.

Para testar, rode: `go run exercicio/main.go`
