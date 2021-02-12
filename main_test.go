package main

import "testing"

func TestSoma(t *testing.T) {
	resultado := soma(5, 5)
	esperado := 10

	if resultado != esperado {
		t.Errorf("soma(5, 5) \n resultado: %v \n esperado: %v ", resultado, esperado)
	}	
}
