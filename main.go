package main

import (
	"fmt"
	// "github.com/google/uuid"
	"net/http"
	// "math"
)

// var a string // fortemente tipado

func main() {
	// var (
	// 	NameSpaceDNS  = uuid.Must(uuid.Parse("6ba7b810-9dad-11d1-80b4-00c04fd430c8"))
	// 	Nil           uuid.UUID // empty UUID, all zeros
	// )

	res, err := http.Get("https://pokeapi.co/api/v2/pokemon/ditto");
	fmt.Println(res.Body)
	fmt.Println("================================================================")
	fmt.Println(err)

	// fmt.Println(math.Soma(1, 2))
	// fmt.Println(err)

	// resultado := Soma(1, 2)
	// a := "Guilherme"
	// b := 10
	// c := 3.144
	// d := false
	// e := `opa

	// bão?`

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)

	// fmt.Printf("%T \n", a)

	// Person := uuid.RFC4122
	// fmt.Println(Nil)
	// fmt.Println(NameSpaceDNS)

	// fmt.Println(resultado)
}

func Soma(a int, b int) int {
	return a + b
}