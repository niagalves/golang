package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo;
}

func main() {
	x, y := trocar(2, 77)
	fmt.Println(x, y)
}
