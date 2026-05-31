package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// 1. Struct e Struct Tags
type Usuario struct {
	ID    int    `json:"id"`            // tag: na conversão para JSON, use chave 'id'
	Nome  string `json:"name"`          // tag: use 'name'
	Email string `json:"email,omitempty"` // omitempty: se estiver vazio (""), ignore este campo no JSON
}

// 2. Métodos e Ponteiros
// Receiver por VALOR (recebe uma cópia)
func (u Usuario) ImprimirDados() {
	fmt.Printf("ID: %d | Nome: %s\n", u.ID, u.Nome)
}

// Receiver por PONTEIRO (consegue modificar a struct original)
func (u *Usuario) AlterarNome(novoNome string) {
	u.Nome = novoNome
}

// 3. Interfaces (Duck Typing)
// Qualquer tipo que tenha o método Pagar(float64) string será um Pagador
type Pagador interface {
	Pagar(valor float64) string
}

type CartaoCredito struct {
	Numero string
}

// CartaoCredito implementa Pagador implicitamente
func (c CartaoCredito) Pagar(valor float64) string {
	return fmt.Sprintf("Pagamento de R$%.2f aprovado no Cartão %s", valor, c.Numero)
}

type Pix struct {
	Chave string
}

// Pix implementa Pagador implicitamente
func (p Pix) Pagar(valor float64) string {
	return fmt.Sprintf("Pagamento de R$%.2f via Pix para a chave %s", valor, p.Chave)
}

// Função que aceita a Interface Pagador
func processarCheckout(p Pagador, total float64) {
	fmt.Println("[Processando Pagamento]", p.Pagar(total))
}

// 4. Erros Customizados e Empacotamento
var ErrSaldoInsuficiente = errors.New("saldo insuficiente na conta")

func sacar(saldo, valor float64) (float64, error) {
	if valor > saldo {
		// Empacotando o erro raiz com mais contexto (%w)
		return 0, fmt.Errorf("tentativa de saque de %.2f: %w", valor, ErrSaldoInsuficiente)
	}
	return saldo - valor, nil
}

func main() {
	fmt.Println("=== Fase 5: Structs, Métodos, Interfaces e Erros ===\n")

	// 1. Instanciando e Serializando a Struct
	fmt.Println("[Structs e Tags JSON]")
	user := Usuario{ID: 1, Nome: "Guilherme"} // Email está vazio
	
	jsonBytes, err := json.Marshal(user)
	if err == nil {
		fmt.Printf("Struct para JSON: %s\n", string(jsonBytes))
	}

	// 2. Chamando Métodos
	fmt.Println("\n[Métodos: Ponteiro vs Valor]")
	user.ImprimirDados() // Método que lê
	
	user.AlterarNome("Guilherme Atualizado") // Método que altera (usa ponteiro por debaixo dos panos: (&user).AlterarNome())
	fmt.Println("Após AlterarNome:")
	user.ImprimirDados()

	// 3. Polimorfismo e Interfaces Implícitas
	fmt.Println("\n[Interfaces Implícitas]")
	cartao := CartaoCredito{Numero: "****-1234"}
	meuPix := Pix{Chave: "contato@email.com"}

	processarCheckout(cartao, 99.90) // Passando a Struct que age como Pagador
	processarCheckout(meuPix, 50.00)

	// 4. Tratamento e Wrapping de Erros
	fmt.Println("\n[Tratamento e Wrapping de Erros]")
	novoSaldo, err := sacar(100.0, 150.0)
	if err != nil {
		fmt.Println("Erro retornado:", err)
		
		// Verificando se o erro embrulhado (%w) é o ErrSaldoInsuficiente
		if errors.Is(err, ErrSaldoInsuficiente) {
			fmt.Println("-> Detalhe: Ocorreu o erro específico de saldo!")
		}
	} else {
		fmt.Println("Saque concluído. Novo saldo:", novoSaldo)
	}
}
