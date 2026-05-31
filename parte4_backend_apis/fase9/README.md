# Fase 9: Comunicação de Alta Performance (gRPC)

Quando construímos arquiteturas de microserviços, a comunicação entre eles via HTTP/REST com JSON começa a se tornar um gargalo de performance e manutenção. JSON é legível para humanos, mas pesado para máquinas interpretarem em altíssimo volume.

O **gRPC** é um framework RPC (Remote Procedure Call) de alta performance criado pelo Google. Ele resolve esses problemas utilizando duas tecnologias chave:
1. **HTTP/2**: Permite multiplexação (várias requisições na mesma conexão TCP), streaming bidirecional e compressão de headers.
2. **Protocol Buffers (Protobuf)**: O formato de serialização binária.

## Protocol Buffers e o Arquivo `.proto`

Em REST, você geralmente escreve o código e depois documenta a API usando Swagger/OpenAPI.
No gRPC, é o modelo "API-First". Você escreve um arquivo `.proto` independente de linguagem, que atua como o contrato universal do microserviço.

```protobuf
syntax = "proto3";

service UsuarioService {
  rpc ObterUsuario (UsuarioRequest) returns (UsuarioResponse);
}

message UsuarioRequest {
  int32 id = 1;
}

message UsuarioResponse {
  string nome = 1;
  string email = 2;
}
```

## Geração de Código (`protoc`)

Após escrever o `.proto`, você utiliza o compilador `protoc` para gerar automaticamente as classes e interfaces em Go (ou Java, Python, Node, etc). Isso elimina a necessidade de escrever clientes HTTP na mão e garante que o cliente e o servidor estejam sempre com o mesmo "contrato" de tipagem estrita (você não passa uma string onde deveria ser um int32, pois o código nem compila).

## Tipos de Streaming

Além do tradicional Requisição -> Resposta (Unary), o gRPC brilha no streaming:
- **Server Streaming**: O cliente faz um pedido e recebe um fluxo contínuo de dados (ex: feed de ações da bolsa).
- **Client Streaming**: O cliente envia um fluxo contínuo e o servidor responde no final (ex: upload de arquivo grande).
- **Bidirectional Streaming**: Cliente e servidor enviam e recebem dados de forma totalmente assíncrona pela mesma conexão.

---

### Executando o exemplo prático

Como o gRPC exige a instalação do compilador `protoc` e a geração de arquivos, o arquivo `main.go` desta fase traz uma *simulação didática* de como o código final (gerado + sua lógica) fica estruturado em Go.
Abra e rode para entender o padrão de métodos RPC:
```bash
go run main.go
```

## 🚀 Exercício Prático

Agora é a sua vez! Entre na pasta `exercicio/` e abra o arquivo `main.go`.
Seu desafio é:
1. Criar structs simulando o Request/Response gerados pelo Protobuf.
2. Implementar a Interface criada usando a sua própria struct simulando um servidor RPC, e fazer uma requisição no método main.

Para testar, rode: `go run exercicio/main.go`
