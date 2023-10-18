package intro

import (
	"fmt"
	"reflect"
)

type Persona struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre"`
	Correo string `json:"correo"`
	Edad   int    `json:"edad"`
}

func Struct() {

	manuel := Persona{
		Id:     1,
		Nombre: "manuel roa",
		Correo: "manuel@mail.com",
		Edad:   32,
	}

	fmt.Printf("%+v \n", manuel)
	fmt.Println(manuel)
	manuel.Nombre = "manuel alejandro roa ojeda"

	fmt.Println(manuel.Nombre)

	sofia := new(Persona)
	fmt.Println(reflect.TypeOf(manuel))
	fmt.Println(reflect.TypeOf(sofia))
	fmt.Println(sofia)
}
