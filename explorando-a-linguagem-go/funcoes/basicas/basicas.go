package main

import "fmt"

func f1() {
	fmt.Println("Função 1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("%s %s \n", p1, p2)
}

func f3(p string) string {
	return p
}

func main() {
	f1()
	f2("teste", "teste2")
	fmt.Println(f3("nome"))
}
