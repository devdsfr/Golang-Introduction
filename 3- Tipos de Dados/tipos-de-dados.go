package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = -100000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// alias
	// int32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	// byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)
	// fim dos inteiros

	// numeros reais
	var numeroReal float32 = 123.45
	fmt.Println(numeroReal)

	var numeroReal2 float64 = 1230000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)
	// fim numeors reais

	// strings
	var str string = " Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)
	// fim strings

	texto := 5
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
