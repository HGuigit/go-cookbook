# Fase 10: Banco de Dados e Transações

O Go possui uma abordagem agnóstica para bancos de dados através do pacote built-in `database/sql`.
Ele não é um ORM (Object-Relational Mapper) pesado como o Hibernate do Java ou o Entity do C#. Ele é uma interface limpa para executar queries SQL cruas, com suporte massivo a concorrência e pooling de conexões embutido.

## Drivers e `database/sql`

Para se conectar a um banco (como PostgreSQL, MySQL ou SQLite), você importa um **Driver** (geralmente usando um underscore `_` no import para registrar o driver silenciosamente) e usa a API do `database/sql` para conversar com ele.
Para o PostgreSQL, por exemplo, o driver mais moderno e recomendado é o `pgx`.

Para casos onde mapeamento de struct para banco de dados é muito repetitivo, a comunidade prefere ferramentas como o `sqlx` (que facilita o scan de colunas para structs sem esconder o SQL) em vez de ORMs densos como o `GORM`.

## Connection Pooling (Pool de Conexões)

Você **NÃO** deve abrir e fechar a conexão do banco a cada requisição HTTP.
A chamada `sql.Open()` retorna um pool de conexões gerenciado internamente pelo Go. Ele cria e destrói conexões TCP com o banco dinamicamente, mantendo conexões ociosas abertas para alta performance.
Você configura esse pool com `SetMaxOpenConns()` e `SetMaxIdleConns()`.

## Transações (ACID) atreladas ao Context

Operações que alteram dados em múltiplas tabelas (como transferir dinheiro de A para B) precisam ser atômicas. O Go usa `db.BeginTx(ctx, options)` para iniciar transações.
O uso de **Context** é essencial aqui: se o request HTTP for cancelado pelo usuário no meio de uma transação demorada de banco de dados, o Go cancela a query no banco instantaneamente, liberando os Locks e economizando recursos do Servidor de Banco.

---

### Executando o exemplo prático

Para não exigir a instalação de um banco Postgres na sua máquina agora, o exemplo utiliza uma simulação do comportamento do pacote `database/sql`. Abra e rode:
```bash
go run main.go
```

## 🚀 Exercício Prático

Agora é a sua vez! Entre na pasta `exercicio/` e abra o arquivo `main.go`.
Seu desafio é:
1. Criar um contexto que expira intencionalmente mais rápido do que a lentidão do nosso Banco Simulador.
2. Executar a Query e validar se você consegue receber e tratar o erro de `contexto expirado`. 

Para testar, rode: `go run exercicio/main.go`
