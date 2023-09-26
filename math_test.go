package main

import "testing"

func TestMath(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado é inválido! Resultado obtido %d. Esperado %d", total, 30)
	}
}
