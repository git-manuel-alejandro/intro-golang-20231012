package intro

import (
	"fmt"
	"reflect"
)

func Recursividad(n int) {
	fmt.Println("valor de n : ", n)

	if n > 0 {
		n--
		Recursividad(n)
	}

}

func GoRoutines() {
	canal := make(chan string)

	go func() {
		canal <- Imprime("manuel")
	}()

	fmt.Println(<-canal)

}

func Imprime(nombre string) string {
	return "hola " + nombre
}

func Clousures(num int) func() int {
	var valor = num
	var i = 0

	return func() int {
		i++
		return valor * i
	}

}

func MultipelReturn(nombre, apellido string, edad int) (string, string, int) {

	return nombre, apellido, edad

}
func Slices() {
	var slices = make([]int, 1)

	for i := 0; i < 99; i++ {
		slices = append(slices, (i+1)*2)
	}

	fmt.Println(slices)
	fmt.Println(len(slices))
}

func Arreglos() {
	var arr [8]int
	arr[1] = 1

	for i := 0; i < len(arr); i++ {
		arr[i] = i * 2
	}

	var otro = &arr
	fmt.Println(arr)
	otro[0] = 10
	fmt.Println(*otro)

	fmt.Println("dirección de memoria : ", &arr)
	fmt.Println("dirección de memoria : ", *otro)
	fmt.Println("==================================================")

	var otro2 [100]int

	for i := 0; i < 100; i++ {
		otro2[i] = (i + 1) * 3
	}
	fmt.Println(otro2)
	fmt.Println(reflect.TypeOf(otro2))

}

func Iteraciones(n int) {
	// for i := 0; i <= n; i++ {
	// 	fmt.Printf("valor de n es : %d \n", i)
	// }

	for i := 10; i >= 0; i-- {
		fmt.Printf("valor de n es : %d \n", i)
	}
}

func VariableCondicion() {

	if edad := 27; edad >= 18 {
		fmt.Println("mayor de edad")
		fmt.Println(edad)
		return
	}

	fmt.Println("menor de edad")
}

func OperadoresLogicos(edad int, pais string) {
	if edad >= 18 && pais == "chile" {
		fmt.Println("you are ready for war")
		return
	}

	fmt.Println("not ready for war")
}

func Condicionales(edad int) {

	if edad >= 18 {
		fmt.Println("mayor de edad")

		return
	}

	fmt.Println("menor de edad")

}

func Punteros() {
	var txt = "texto"
	fmt.Println("direccion : ", &txt)

	fmt.Println("reflect type of : ", reflect.TypeOf(txt))

	var puntero = &txt
	fmt.Println("valor :", *puntero)
}

func TypeOf() {
	var txt = "hola"
	var num = 1

	fmt.Println(reflect.TypeOf(txt))
	fmt.Println(reflect.TypeOf(num))
}
