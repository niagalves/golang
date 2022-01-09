package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"João":     11343.44,
		"Gabriela": 65555.99,
		"Pedro":    650.00,
	}

	funcsESalarios["Rafael Filho"] = 132.0

	delete(funcsESalarios, "João")
	fmt.Println(funcsESalarios)

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
