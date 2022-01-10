package main

import "fmt"

func imprimeAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados")
	for i, aprovados := range aprovados {
		fmt.Printf("%d) %s \n", i+1, aprovados)
	}
}

func main() {
	aprovados := []string{"maria", "pedro", "rafael"}

	imprimeAprovados(aprovados...)
}
