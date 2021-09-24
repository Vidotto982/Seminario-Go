package main

import (
	"testing"
	// "fmt"

	"github.com/stretchr/testify/assert"
	"tudai2021.com/model"
)

func TestChangeLength(t *testing.T){
	p := model.NewResult("TX02AB")
	changeLength(&p, 0002)
	assert.Equal(t, p.Length,0002, "los valores no coinciden")

}
func TestChangeType(t *testing.T){
	p := model.NewResult("TX02AB")
	changeType(&p, "TX")
	assert.Equal(t, p.Type,"TX", "los valores no coinciden")
}
func TestChangeValue(t *testing.T){
	p := model.NewResult("TX02AB")
	changeValue(&p, "05")
	assert.Equal(t, p.Value,"05", "los valores no coinciden")
	
}

func TestValidateType(t *testing.T) {
	err := validateType("NN","123")
	err1 := validateType("NN","ABC")
	err2 := validateType("XX","ABC")

	assert.Equal(t, err , nil,"El Tipo del valor no es valido" )
	assert.Equal(t, err1 , nil,"El Tipo del valor no es valido" )
	assert.Equal(t, err2 , nil,"El Tipo del valor no es valido" )

}
func TestValidateLength(t *testing.T) {
	err := validateLength(4,"123")
	err1 := validateLength(3,"123")
	assert.Equal(t, err , true, "El largo del valor no es valido")
	assert.Equal(t, err1 , true, "El largo del valor no es valido")
	
}
