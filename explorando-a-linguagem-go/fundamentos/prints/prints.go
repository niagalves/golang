package main

import "fmt"

func main() {
	fmt.Print("Print")   // mesma linha
	fmt.Print("Print 2") // mesma linha

	fmt.Println("Nova") // quebra linha

	x := 3.141516

	xs := fmt.Sprint(x)

	fmt.Println("O valor de x é" + xs)
	fmt.Printf("O valor de x é %.2f", x) // renderiza apenas duas casas decimais.

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
