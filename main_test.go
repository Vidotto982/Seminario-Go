package main

import (
	"testing"

	"tudai2021.com/model"
)

func TestEstructura(t *testing.T) {
	p := model.NewResult(2, "4213")
	Estructura(&p, "asd")

	if p.Length < "0" && p.Length > "9" {
		t.Error("No es un codigo valido")

	}

}
