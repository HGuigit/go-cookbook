package main

import "fmt"

// =========================================================
// CAMADA 1: ENTITIES (Regras e Estruturas Core)
// Na vida real: internal/entity/usuario.go
// =========================================================
type Usuario struct {
	ID    int
	Nome  string
	Ativo bool
}

// =========================================================
// CAMADA 2: REPOSITÓRIOS (Adapters de Saída - Banco de Dados)
// Na vida real: internal/repository/usuario_repository.go
// =========================================================

// A Interface é a "Porta". Fica no nível de negócio.
type UsuarioRepository interface {
	Salvar(u Usuario) error
}

// A implementação real (o "Adapter" do Postgres)
type PostgresRepository struct {
	// dbConn *sql.DB
}

func (r *PostgresRepository) Salvar(u Usuario) error {
	fmt.Printf("[Postgres] Insert INTO usuarios (nome) VALUES (%s)\n", u.Nome)
	return nil
}


// =========================================================
// CAMADA 3: SERVICES (Usecases / Lógica de Negócio)
// Na vida real: internal/service/usuario_service.go
// =========================================================

type UsuarioService struct {
	// A Mágica: O service DEPENDE da Interface, não do Postgres.
	// Isso permite injetar MockRepository nos testes!
	repo UsuarioRepository 
}

// Construtor idiomático
func NewUsuarioService(r UsuarioRepository) *UsuarioService {
	return &UsuarioService{repo: r}
}

// A Regra de Negócio Pura
func (s *UsuarioService) CadastrarUsuario(nome string) error {
	novoUsuario := Usuario{
		ID:    99,
		Nome:  nome,
		Ativo: true,
	}

	// Salva no banco e retorna
	return s.repo.Salvar(novoUsuario)
}

// =========================================================
// CAMADA 4: HANDLERS (Adapters de Entrada - HTTP/API)
// Na vida real: internal/handler/usuario_handler.go
// =========================================================

type UsuarioHandler struct {
	svc *UsuarioService // O handler depende do service
}

func (h *UsuarioHandler) HandleCadastroPOST(payloadNome string) {
	fmt.Printf("[HTTP POST] Recebido payload: %s\n", payloadNome)
	
	// Repassa a responsabilidade pro UseCase
	h.svc.CadastrarUsuario(payloadNome)
	
	fmt.Println("[HTTP POST] Resposta HTTP 201 Created retornada.")
}


// =========================================================
// ENTRYPOINT (Wire up)
// Na vida real: cmd/api/main.go
// =========================================================
func main() {
	fmt.Println("=== Fase 14: Estrutura Clean Architecture e Injeção de Dependência ===\n")

	// 1. Iniciamos as dependências de infraestrutura (Banco, Redis)
	dbRepo := &PostgresRepository{}

	// 2. Injetamos o Banco no Service (Regra de negócio)
	service := NewUsuarioService(dbRepo)

	// 3. Injetamos o Service no Handler (Controlador HTTP)
	handler := &UsuarioHandler{svc: service}

	// 4. Simulando uma request HTTP batendo no Controller
	handler.HandleCadastroPOST("Guilherme - Go Developer")
}
