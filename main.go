//fmt.Println(country[0:2] + " " + country[2:4] + " " + country[5:])
package main

import (
	"fmt"
	"strconv"
)

//"strings"

type Result struct {
	Type   string
	Value  string
	Length int
}
type ResultInterface interface {
	structResult(tipo string, valor string, length int) (*Result, error)
}

func main() {
	country := "XK0300007"
	length, _ := strconv.Atoi(country[5:])
	result := Result{country[0:2], country[2:4], length}
	fmt.Println(result)
}
