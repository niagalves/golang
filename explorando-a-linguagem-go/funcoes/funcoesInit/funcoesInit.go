package main

import "fmt"

// função por padrão é executada por primeiro.
// utilziado quando precisa de algumas funções principais serem iniciadas antes das demais
// funções do projeto.

func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}
