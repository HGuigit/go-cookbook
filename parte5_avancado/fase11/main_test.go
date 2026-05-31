package main

import "testing"

// =========================================================
// O CÓDIGO DE TESTES UNITÁRIOS (Table-Driven Tests)
// =========================================================

func TestClassificaIdade(t *testing.T) {
	// 1. A Tabela de Cenários (Struct Anônima)
	cenarios := []struct {
		nomeDoTeste string
		input       int
		esperado    string
		querErro    bool
	}{
		{"Idade de uma Criança", 5, "Criança", false},
		{"Idade de um Adolescente", 15, "Adolescente", false},
		{"Idade de um Adulto", 30, "Adulto", false},
		{"Idade de um Idoso", 65, "Idoso", false},
		{"Erro ao passar idade negativa", -5, "", true},
	}

	// 2. O laço que itera sobre a tabela rodando os sub-testes
	for _, cenario := range cenarios {
		t.Run(cenario.nomeDoTeste, func(t *testing.T) { // t.Run isola cada teste
			
			resultado, err := ClassificaIdade(cenario.input)

			// Validando erro
			if cenario.querErro && err == nil {
				t.Errorf("Esperava um erro, mas não retornou nada")
			}
			if !cenario.querErro && err != nil {
				t.Errorf("Não esperava erro, mas ocorreu: %v", err)
			}

			// Validando resultado
			if resultado != cenario.esperado {
				// %q formata strings com aspas, bom para visualizar
				t.Errorf("Esperado %q, mas obtido %q", cenario.esperado, resultado)
			}
		})
	}
}

// =========================================================
// O CÓDIGO DE BENCHMARK (Medição de Performance)
// =========================================================

// Benchmark funções SEMPRE começam com "Benchmark" e recebem *testing.B
func BenchmarkMontarStringIneficiente(b *testing.B) {
	// b.N é o número de repetições. O runtime do Go ajusta b.N dinamicamente
	// até que a medição seja estatisticamente relevante (geralmente dura 1s por default)
	for i := 0; i < b.N; i++ {
		MontarStringIneficiente(100)
	}
}
