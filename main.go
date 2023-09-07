package main

import (
	"fmt"

	"github.com/jlromgaz/godesde0/variables"
)

func main() {
	variables.MuestroEnteros()
	variables.RestoVariables()

	estado, texto := variables.ConviertoaTexto(1234)
	fmt.Println("estado = ", estado)
	fmt.Println("texto = ", texto)
}
