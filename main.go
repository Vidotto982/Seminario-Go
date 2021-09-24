//fmt.Println(country[0:2] + " " + country[2:4] + " " + country[5:])
package main

import (
	"tudai2021.com/model"
	"strconv"
	"errors"
	"unicode"
	"fmt"
)


func changeType(p *model.Result, tipo string) {
	p.Type = tipo
	//modificamos el tipo de una cadena de texto creada

}
func changeValue(p *model.Result, valor string) {
	p.Value = valor
	//modificamos el valor de una cadena de texto creada
	
}
func changeLength(p *model.Result, largo int) {
	p.Length = largo
	//modificamos el largo de una cadena de texto creada
}
func validateType(tipo string,valor string) error{
	if tipo == "TX" {
		return validateLetterValue(valor)
		//validamos que no tenga numeros el valor definido en la cadena
	}
	if tipo == "NN" {
		return validateValueNumber(valor)
		//validamos que no tenga letras el valor definido en la cadena

	}
	return errors.New("Tipo de valor invalido")

}
func validateLetterValue(valor string) error {
	for _, aux := range valor {
		if unicode.IsDigit(aux) {
			return errors.New("El valor no es valido, contiene numeros")
			//verificamos en este metodo que el valor no contenga ningun numero
			
		}
	}
	return nil
}

func validateValueNumber(valor string) error {
	_, err := strconv.Atoi(valor)
	if err != nil {
		return errors.New("El valor no es valido, contiene letras")
		//verificamos en este metodo que el valor no contenga ninguna letra
	}
	return nil
}
func validateLength(largo int,valor string) bool{
	if	len(valor) == largo{
		return true
		//verificamos en este metodo que el largo del valor sea correcto

	}
	return false
}
func main() {
	p := model.NewResult("TX040001")
	fmt.Println(p)
}
