package main

import "testing"

func TestSoma(t *testing.T) {

	total := soma(16, 15)

	if total != 31 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 30)
	}

}
