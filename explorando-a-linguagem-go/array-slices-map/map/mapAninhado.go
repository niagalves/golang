package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela": 21212121,
			"Guga":     2121111,
		},
		"N": {
			"Nadia":  221,
			"Nairon": 140,
		},
	}

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
