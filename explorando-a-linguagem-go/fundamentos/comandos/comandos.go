package main

import "fmt"

func main() {
	fmt.Printf("Outro programa em!")
}

// passando go help e o comando que você quer ele trás
// as informações necessárias sobre aquele comando

// Você pode ter acesso a documentação do golang sem ter acesso a internet utilizando o comando
// godoc -http=:6060

// Mostra informações sobre o ambiente
// go env

// faz uma verificação do código atrás de erros
// go vet nome-arquivo.go

// faz o build do projeto
// go build

// faz o build e já executa
// go run nome-arquivo.go
