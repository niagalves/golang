package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 4

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Módulo, resto da divisão =>", a%b)

	// bitwise
	fmt.Println("E =>", a&b)
	fmt.Println("Ou =>", a|b)
	fmt.Println("Xor =>", a^b)

	c := 3.0
	d := 2.0

	// outras operações usando math..
	fmt.Println("Maior =>", math.Max(float64(c), float64(b)))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponenciação =>", math.Pow(c, d))
}
