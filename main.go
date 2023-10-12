package main

import (
	"curso/intro"
	"fmt"
)

func main() {
	nombre, apellido, edad := intro.MultipelReturn("manuel", "roa", 32)
	fmt.Printf("%s %s tiene %d años de edad", nombre, apellido, edad)
	// intro.Slices()
	// intro.Arreglos()
	// intro.Iteraciones(10)
	// intro.VariableCondicion()
	// intro.OperadoresLogicos(2, "chile")
	// intro.Condicionales(19)
	// intro.Punteros()
	// intro.TypeOf()
}
