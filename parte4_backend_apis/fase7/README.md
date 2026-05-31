# Fase 7: Servidores HTTP e Middlewares

O Go possui uma biblioteca padrão excepcionalmente poderosa para redes e HTTP. Na grande maioria dos projetos, você **não precisa de frameworks** pesados como Spring ou Express. O pacote nativo `net/http` é otimizado para produção e lida com milhares de requisições concorrentes naturalmente (ele abre uma Goroutine automaticamente por requisição).

## A Biblioteca `net/http` e o ServeMux

Para iniciar um servidor HTTP, os passos básicos são:
1. Criar um **Handler** (uma função que recebe a Request e escreve a Response).
2. Associar esse Handler a uma rota usando um **Router/Mux** (Multiplexador).
3. Iniciar o servidor escutando em uma porta específica.

No Go 1.22+, o roteador padrão (`http.ServeMux`) ganhou suporte nativo para parâmetros de rota (ex: `/usuarios/{id}`) e restrição de verbos HTTP (`GET /usuarios`), tornando roteadores de terceiros muitas vezes desnecessários.

## Serialização e Desserialização JSON

Como APIs REST geralmente trafegam JSON, usamos o pacote padrão `encoding/json`.
- Para receber dados (POST/PUT): usamos `json.NewDecoder(r.Body).Decode(&struct)` para transformar o JSON em uma Struct do Go.
- Para enviar dados (GET/Respostas): usamos `json.NewEncoder(w).Encode(struct)` para converter nossa Struct do Go de volta para a rede.

## Middlewares (Decorators)

Um Middleware é uma função que "envolve" o Handler principal de uma rota, executando código antes e/ou depois da lógica de negócio. 
Eles são úteis para:
- Loggar as requisições (tempo de duração, IP).
- Verificar autenticação/autorização (Validar JWT).
- Tratamento global de *panics* (Recovery).

Em Go, criamos middlewares escrevendo funções que recebem um `http.Handler` e retornam um `http.Handler`.

---

### Executando o exemplo prático

Vamos subir o servidor! Abra o terminal nesta pasta e rode:
```bash
go run main.go
```
E acesse as rotas no navegador ou Postman:
- `GET http://localhost:8080/ping`
- `POST http://localhost:8080/api/produtos` (Envie um JSON no Body: `{"nome": "Teclado", "preco": 199.90}`)

## 🚀 Exercício Prático

Agora é a sua vez! Entre na pasta `exercicio/` e abra o arquivo `main.go`.
Seu desafio é:
1. Criar um Handler HTTP customizado que responda com uma saudação.
2. Registrar esse Handler no roteador padrão em uma porta diferente (`:8081`).

Para testar, rode: `go run exercicio/main.go` e acesse no navegador.
