package main

import "fmt"

// defer é uma função que atrasa a execução de determinada setenção de código.
// muito comum por exemplo abrir uma conexão de banco de dados e quer ser lembrado de fechar a conexão.

func obterValorAprovado(numero int) int {
	defer fmt.Println("Fim!")

	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	} else {
		fmt.Println("Valor baixo...")
		return numero
	}
}

func main() {
	obterValorAprovado(50400)
}
