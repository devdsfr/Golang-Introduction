package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Print("Herança")

	p1 := pessoa{"Daniel", "Ramos", 30, 182}
	fmt.Print(p1)

	p2 := estudante{p1, "Engenharia", "USP"}
	fmt.Print(p2)

}
