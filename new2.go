package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre   string
	Apellido string
	Frases   []string
}

func main() {
	p1 := persona{
		Nombre:   "James",
		Apellido: "Bond",
		Frases:   []string{"Shaken, not stirred", "¿Algún último deseo?", "Nunca digas nunca."},
	}

	bs, err := aJSON(p1)

	fmt.Println(string(bs))

}

// aJSON también necesita retorna un error
func aJSON(a interface{}) []byte {
	bs, _ := json.Marshal(a)
}
