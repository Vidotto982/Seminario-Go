//fmt.Println(country[0:2] + " " + country[2:4] + " " + country[5:])
package main

import (
	"fmt"
	"strconv"

	"tudai2021.com/model"
)

//"strings"

func Estructura() {
	country := "XK0300007"
	length, _ := strconv.Atoi(country[5:])
	result := model.Result{country[0:2], country[2:4], length}
	fmt.Println(result)
}

func main() {
	Estructura()
}
