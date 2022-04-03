package main

import "fmt"

type usuario struct {
	nome     string
	idade    int
	endereco endereco
}

type endereco struct {
	rua    string
	numero int
}

func main() {
	fmt.Print("Arquivos structs")

	var u usuario
	u.nome = "Daniel"
	u.idade = 5
	fmt.Println(u)

	enderecoCompleto := endereco{"Rua azul", 5464}

	usuario2 := usuario{"Daniel", 5, enderecoCompleto}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Daniel"}
	fmt.Println(usuario3)

}
