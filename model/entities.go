package model
import (
	"strconv"
)
type Result struct {
	Type   string
	Length int
	Value  string
}

func NewResult(cadena string) Result{
	length, _ := strconv.Atoi(cadena[2:4])
	//convertimos los datos que estan indicados en las posiciones 2y3 de la cadena en int
	return Result{cadena[0:2], length,cadena[4:]}
}
