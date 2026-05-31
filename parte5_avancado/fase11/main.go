package main

import "errors"

// =========================================================
// O CÓDIGO DA APLICAÇÃO (A ser testado)
// =========================================================

// ClassificaIdade classifica a idade em categorias ou retorna erro se for inválida
func ClassificaIdade(idade int) (string, error) {
	if idade < 0 {
		return "", errors.New("idade não pode ser negativa")
	}

	switch {
	case idade < 13:
		return "Criança", nil
	case idade < 18:
		return "Adolescente", nil
	case idade < 60:
		return "Adulto", nil
	default:
		return "Idoso", nil
	}
}

// Para Benchmark: Uma função que concatena strings ineficientemente
func MontarStringIneficiente(qtd int) string {
	res := ""
	for i := 0; i < qtd; i++ {
		res += "A" // Isso gera muita alocação de memória na heap
	}
	return res
}
